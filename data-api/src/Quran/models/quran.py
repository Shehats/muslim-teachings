from models.base_model import (BaseModel, Base)
from sqlalchemy import (Column)

from sqlalchemy.orm import (relationship)

class QuranModel(Base):
    quran_id = Column(Integer, primary_key=True)
    surahs = relationship('Surah')