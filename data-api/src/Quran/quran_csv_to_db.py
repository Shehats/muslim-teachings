import numpy as np
import pandas as pd
import os
from os import walk

from constants.constants import (
    SURA_ID,
    VERSE_ID,
    VERSE_ID,
    AYAH_TEXT,
    QURAN_SURAHS
)
from random import random
from typing import List
from models.ayah import Ayah
from models.surah import Surah
import time

from google.cloud import speech_v1
from google.cloud.speech_v1 import enums

import base64
import json
from collections import deque

def maybe_create_dir(path, dir_name=None):
    dir_path = "{}/{}".format(path, dir_name) if dir_name else path
    if not os.path.exists(dir_path):
        os.mkdir(dir_path)

def maybe_move_file(path, new_path, old_file, new_file):
    old_name = "{}/{}".format(path, old_file)
    new_name = "{}/{}".format(path, new_file)
    os.rename(old_name, new_name)
    target_name = "{}/{}".format(new_path, new_file)
    os.replace(new_name, target_name)


class QuranCSVToDB:
    def __init__(self, Quran_path):
        self._text_data = []
        self._voice_data = []
        self._Quran_path = Quran_path
        self._data = None
        self._surahs_map = {surah['id']: surah for surah in QURAN_SURAHS}
        self._audio_files = {}
        self._Quran = [None]
        self._total_ayat_count = 0
        self._recitations = {}
        self._lazy_reversed_surah = {surah['english_name']: surah['ayah_count'] for surah in QURAN_SURAHS}
        

    async def read_text_data(self):
        if self._Quran[len(self._Quran) - 1] is not None:
            return
        self._Quran = await self._read_from_data()

    async def _read_from_data(self):
        data = pd.read_csv(self._Quran_path, ",")
        surah: Surah = None
        quran = []
        stack = [(i, row) for i, row in data.iterrows()]
        while stack:
            _, row = stack[len(stack)-1]
            surah_id = int(row['SuraID'])
            ayat = []
            while stack and int(row['SuraID']) == surah_id:
                ayah_count, row = stack.pop()
                surah_id = row['SuraID']
                curn_surah = self._surahs_map[surah_id]
                ayah: Ayah = Ayah()
                ayah.AyahText = row['AyahText']
                ayah.AyahNumber = int(row['VerseID'])
                ayah.AyahRelativeNumber = ayah_count
                ayat.insert(0,ayah)
                if stack:
                    _, row = stack[len(stack)-1]
            self._total_ayat_count += len(ayat)
            
            surah: Surah = Surah()
            surah.ArabicName = curn_surah['arabic_name']
            surah.EnglishName = curn_surah['english_name']
            surah.Gozoa = curn_surah['gozoa']
            surah.AyatCount = curn_surah['ayah_count']
            surah.Ayat = ayat
            quran.append(surah)

        quran = [None] + quran[::-1]
        return quran

    async def walk_audio_files(self, path, reciter_name, split_str=""):
        path = "{}/{}".format(path, reciter_name)
        self._audio_files[reciter_name] = await self._walk_audio_files(path, split_str)

    async def _walk_audio_files(self, path, split_str=""):
        audio_files = [None]*115
        async with walk(path) as p:
            for (_1, _2, filenames) in p:
                for filename in filenames:
                    name_split = filename.split(split_str)
                    audio_files[int(name_split[0])] =  filename
        
        return audio_files

    async def check_files_in_resources(self, path):
        ret_files = await self._check_files_in_resources(path)
        return ret_files

    async def _check_files_in_resources(self, path):
        ret_files = deque()
        async with walk(path) as p:
            for (_, dirnames, _) in walk(path):
                for dirname in dirnames:
                    surah_id, surah_name = dirname.split('_')
                    for (_,_,filename) in walk("{}/{}".format(path, dirname)):
                        surah_files = set(filename)

                    if surah_name not in self._lazy_reversed_surah or len(surah_files) != self._lazy_reversed_surah[surah_name]:
                        return None
                    ret_files.append((dirname, surah_id, surah_files))
        return ret_files
                        

    def load_recitations_files_from_file(self, path, reciter_name):
        path = "{}/{}".format(path, reciter_name)
        self.read_text_data()
        data = self.check_files_in_resources(path)
        if not data:
            return False
        
        while data:
            dirname, surah_id, surah_files = data.popleft()
            for ayah in self._Quran[int(surah_id)].Ayat:
                f_name = "{}.mp3".format(ayah.AyahNumber)
                if f_name in surah_files:
                    recitation_file_path = "{}/{}/{}".format(path, dirname,f_name)
                    with open(recitation_file_path, 'rb') as recitation:
                        recitation_base64 = base64.b64encode(recitation.read())
                        ayah.AyahRecitation.append({reciter_name: recitation_base64})
            return True
            

    
    def walk_audio_files_non_ordered_numeric(self, path, reciter_name, split_str=""):
        if reciter_name in self._recitations and len(self._recitations[reciter_name]) == self._total_ayat_count:
            return
        
        self.read_text_data()
        path = "{}/{}".format(path, reciter_name)
        audio_files_num = []
        reversed_file_name = {}
        audio_files_alpha = []
        for (_1, _2, filenames) in walk(path):
            for filename in filenames:
                if '.mp3' in filename:
                    k = filename.split('.')[0]
                    if split_str and split_str in filename:
                        name_split = k.split(split_str)
                        k = name_split[0]
                    if k.isnumeric():
                        audio_files_num.append(int(k))
                        reversed_file_name[str(int(k))] = filename
                    else:
                        audio_files_alpha.append(k)
                        reversed_file_name[k] = filename

        if len(audio_files_num) == self._total_ayat_count:
            self._recitations[reciter_name] = []
            audio_files_num.sort()
            for i, surah in enumerate(self._Quran):
                if surah is not None:
                    ayat = surah.Ayat
                    for ayah in ayat:
                        old_name = reversed_file_name[str(audio_files_num[ayah.AyahRelativeNumber])]
                        new_path = "{}/{}_{}".format(path, i ,surah.EnglishName)
                        entry = ("{}_{}".format(i, surah.EnglishName), new_path, old_name, "{}.mp3".format(ayah.AyahNumber))
                        self._recitations[reciter_name].append(entry)
            
            if reciter_name in self._recitations:
                for surah_name, new_path, old_name, new_name in self._recitations[reciter_name]:
                    maybe_create_dir(path, surah_name)
                    maybe_move_file(path, new_path, old_name, new_name)
