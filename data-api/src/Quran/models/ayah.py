from models.tafseer import (
    Tafseer,
    TafseerModel
)
from models.surah import SurahModel

from typing import List
from models.base_model import (BaseModel, Base)

from sqlalchemy import *
from sqlalchemy.orm import (scoped_session, sessionmaker, relationship,
                            backref)
from sqlalchemy.ext.declarative import declarative_base

class Ayah(BaseModel):
    def __init__(self):
        self._ayah_tafseer: List[Tafseer] = []
        self._ayah_recitation = []

    @property
    def AyahNumber(self)->int:
        return self._ayah_number
    
    @property
    def AyahText(self)->str:
        return self._ayah_text
    
    @property
    def AyahTafseer(self)->List[Tafseer]:
        return self._ayah_tafseer

    @property
    def AyahRelativeNumber(self)->int:
        return self._ayah_relative_number
    
    @property
    def AyahRecitation(self):
        return self._ayah_recitation
    
    @AyahNumber.setter
    def AyahNumber(self, value: int):
        self._ayah_number = value
    
    @AyahText.setter
    def AyahText(self, value: str):
        self._ayah_text = value

    @AyahTafseer.setter
    def AyahTafseer(self, value: List[Tafseer]):
        self._ayah_tafseer = value
    
    @AyahRelativeNumber.setter
    def AyahRelativeNumber(self, value):
        self._ayah_relative_number = value
    
    @AyahRecitation.setter
    def AyahRecitation(self, value):
        self._ayah_recitation = value
    
    def add_recitation(self, value):
        self._ayah_recitation.append(value)
    
    def add_tafseer(self, value):
        self._ayah_tafseer.append(value)

class AyahModel(Base):
    __tablename__ = 'Ayah'
    ayah_id = Column(Integer, primary_key=True)
    ayah_text = Column(String)
    ayah_tafseer = relationship('Tafseer')
    ayah_relative_number = Column(Integer)
    ayah_recitations = Column('AyahRecitation')
    surah = relationship(AyahModel, backref=backref('ayat', uselist=True, cascade='delete,all'))

class AyahRecitation(Base):
    __tablename__ = 'AyahRecitation'
    recitation_id = Column(Integer, primary_key=True)
    reciter_name = Column(String)
    reciter_media = Column(String)
    ayah = relationship(AyahModel, backref=backref('ayah_recitations', uselist=True, cascade='delete,all'))