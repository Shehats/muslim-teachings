from graphene import (ObjectType, relay, List, String, Schema)
from graphene_sqlalchemy import SQLAlchemyConnectionField
from app_connections import (QuranConnection, SurahConnection, AyahConnection, TafseerConnection, SearchResults)
from app_connections import (Ayah, Tafseer, Quran, Surah)
from models.ayah import AyahModel
from models.tafseer import TafseerModel

class AllQuery(ObjectType):
    node = relay.Node.Field()
    quran = SQLAlchemyConnectionField(QuranConnection)
    all_surahs = SQLAlchemyConnectionField(SurahConnection)
    all_ayat = SQLAlchemyConnectionField(AyahConnection)
    all_tafseer = SQLAlchemyConnectionField(TafseerConnection)
    search = List(SearchResults, q=String())

    def resolve_search(self, info, **kwargs):
        q = args.get("q")
        tafseer_query = Tafseer.get_query(info)
        ayah_query = Ayah.get_query(info)

        ayat = ayah_query.filter((AyahModel.ayah_id.contains(q)) | (AyahModel.ayah_text.contains(q))).all()
        tafseers = tafseer_query.filter((TafseerModel.source(q))).all()

        return ayat + tafseers

all_schema = Schema(query=AllQuery, types=[Quran, Surah, Ayah, Tafseer, SearchResults])