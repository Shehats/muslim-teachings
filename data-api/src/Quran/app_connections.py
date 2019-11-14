import graphene
from graphene import relay
from graphene_sqlalchemy import SQLAlchemyObjectType, SQLAlchemyConnectionField
from models.base_model import db_session
from models.ayah import AyahModel
from models.surah import SurahModel
from models.tafseer import TafseerModel
from models.quran import QuranModel

class Quran(SQLAlchemyObjectType):
    class Meta:
        model = QuranModel
        interfaces = (relay.Node, )

class QuranConnection(relay.Connection):
    class Meta:
        node = Quran

class Surah(SQLAlchemyObjectType):
    class Meta:
        model = SurahModel
        interfaces = (relay.Node, )

class SurahConnection(relay.Connection):
    class Meta:
        node = Surah

class Ayah(SQLAlchemyObjectType):
    class Meta:
        model = AyahModel
        interfaces = (relay.Node, )

class AyahConnection(relay.Connection):
    class Meta:
        node = Ayah

class Tafseer(SQLAlchemyObjectType):
    class Meta:
        model = TafseerModel
        interfaces = (relay.Node, )

class TafseerConnection(relay.Connection):
    class Meta:
        node = Tafseer

class SearchResults(graphene.Union):
    class Meta:
        types = (Ayah, Tafseer)
