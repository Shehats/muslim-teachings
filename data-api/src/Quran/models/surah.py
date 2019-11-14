from models.ayah import Ayah
from typing import List
from models.base_model import (BaseModel, Base)
import inspect

from sqlalchemy import (Column, String, Integer)

from sqlalchemy.orm import (relationship, backref)

class Surah(BaseModel):
    def __init__(self):
        self._ayat: List[Ayah] = []

    @property
    def ArabicName(self)->str:
        return self._arabic_name
    
    @property
    def EnglishName(self)->str:
        return self._english_name
    
    @property
    def AyatCount(self)->int:
        return self._ayat_count
    
    @property
    def Ayat(self)->List[Ayah]:
        return self._ayat
    
    @property
    def PageNumber(self)->int:
        return self._page_number

    @property
    def Gozoa(self)->int:
        return self._gozoa
    
    @ArabicName.setter
    def ArabicName(self, value):
        self._arabic_name = value
    
    @EnglishName.setter
    def EnglishName(self, value):
        self._english_name = value
    
    @Ayat.setter
    def Ayat(self, value):
        self._ayat = value
    
    @PageNumber.setter
    def PageNumber(self, value):
        self._page_number = value
    
    @Gozoa.setter
    def Gozoa(self, value):
        self._gozoa = value
    
    @AyatCount.setter
    def AyatCount(self, value):
        self._ayat_count = value
    
    def add_ayah(self, value: Ayah):
        self._ayat.append(value)

class SurahModel(Base):
    __tablename__ = 'Surah'
    surah_id = Column(Integer, primary_key=True)
    arabic_name = Column(String)
    english_name = Column(String)
    ayah_count = Column(Integer)
    ayat = relationship('Ayah')