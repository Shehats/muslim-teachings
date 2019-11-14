from json import dumps
from kafka import KafkaProducer
from .app_constants import (DEFAULT_BOOTSTRAP_SERVERS, DEFAULT_PRODUCER_VALUE_SERIALIZER)

class AppProducer:
    def __init__(self, bootstrap_servers=DEFAULT_BOOTSTRAP_SERVERS, value_serializer=DEFAULT_PRODUCER_VALUE_SERIALIZER):
        self._producer = KafkaProducer(bootstrap_servers=bootstrap_servers, value_serializer=value_serializer)
        