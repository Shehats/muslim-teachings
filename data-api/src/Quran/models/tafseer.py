from typing import List
from models.base_model import BaseModel
from models.base_model import Base

from sqlalchemy import (Column, String, Integer)

from sqlalchemy.orm import (relationship, backref)

class Tafseer(BaseModel):
    def __init__(self):
        self._urls: List[str] = []
    
    @property
    def Source(self)->str:
        return self._source
    
    @property
    def URLS(self)->List[str]:
        return self._urls
    
    @property
    def Content(self)->str:
        return self._content

    @property
    def Audio(self):
        return self._audio
    
    @Source.setter
    def Source(self, value: str):
        self._source = value
    
    @URLS.setter
    def URLS(self, value: List[str]):
        self._urls = value

    @Content.setter
    def Content(self, value: str):
        self._content = value
    
    @Audio.setter
    def Audio(self, value):
        self._audio = value

    def add_url(self, value):
        self._urls.append(value)

class TafseerModel(Base):
    __tablename__ = 'Tafseer'
    id = Column(Integer, primary_key=True, autoincrement=True)
    language = Column(String)
    audio_url = Column(String)
    audio = Column(String)
    source = Column(String)
    tafseer_urls = relationship('Url')

class Url(Base):
    __tablename__ = 'Url'
    id = Column(Integer, primary_key=True, autoincrement=True)
    url = Column(String)
    tafseer = relationship(TafseerModel, backref=backref('urls',
                        uselist=True,
                        cascade='delete,all'))

    