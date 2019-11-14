from json import (dumps, loads)
from sqlalchemy import *
from sqlalchemy.orm import (scoped_session, sessionmaker, relationship,
                            backref)
from sqlalchemy.ext.declarative import declarative_base

engine = create_engine('sqlite:///database.sqlite3', convert_unicode=True)
db_session = scoped_session(sessionmaker(autocommit=False,
                                         autoflush=False,
                                         bind=engine))

Base = declarative_base()
# We will need this for querying
Base.query = db_session.query_property()

class BaseModel(object):
    def jsonify(self):
        return dumps(self.__dict__).encode('utf-8')

    def from_json(self, value):
        return loads(value.decode('utf-8')) 