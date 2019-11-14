from pymongo import MongoClient
from models.base_model import BaseModel

DEFAULT_MONGO_URL, DEFAULT_MONGO_DB = 'localhost:27017', 'Quran_db'
class DBManager:
    def __init__(self, mongo_url: str = DEFAULT_MONGO_URL, db_name = DEFAULT_MONGO_DB):
        self._client = MongoClient("mongodb://{}/".format(mongo_url))
        self._db = self._client[db_name]
    
    def create_document_data(self, data: object):
        data_dict = {}
        for k, v in data.__dict__:
            if isinstance(type(v), object) or issubclass(type(v), object):
                data_dict[k] = self.create_document_data(v)
            else:
                data_dict[k] = v
        return data_dict
                
    def save(self, data: object):
        self._db.insert_one(self.create_document_data(data))
        return True

    def findById(self, data_id):
        return self._db.find_one({'_id': data_id})

    def findOneByField(self, field_name, field_value):
        query_dict = {}
        query_dict[field_name] = field_value
        return self._db.find_one(query_dict)
    
    def findAllByField(self, field_name, field_value):
        query_dict = {}
        query_dict[field_name] = field_value
        return self._db.find(query_dict)
    

    def findAll(self):
        return self._db.find({})

    def update(self, data, data_id = None):
        if not data_id and '_id' not in data:
            all_data = self.findAll()
            obj = data if isinstance(type(data), dict) else data.__dict__
            similarities = 0
            object_to_update = None
            while all_data:
                curn = all_data.pop()
                curn_similar = 0
                for k,v in obj:
                    if k in curn and curn[k] == v:
                        curn_similar += 1
                if curn_similar > similarities:
                    object_to_update = curn
                    similarities = curn_similar
            
            if object_to_update:
                self._db.update_one({'_id': object_to_update.get('_id')}, data)
                return True
            return False
        data_id = data_id if data_id else data.get('_id')
        self._db.update_one({'_id': data_id}, data)
        return True
