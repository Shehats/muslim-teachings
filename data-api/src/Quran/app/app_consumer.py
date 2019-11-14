from .app_constants import (DEFAULT_BOOTSTRAP_SERVERS, 
                            DEFAULT_AUTO_OFFSET_RESET, 
                            DEFAULT_COMMIT_FLAG,
                            DEFAULT_CONSUMER_VALUE_SERIALIZER)
from kafka import KafkaConsumer
class AppConsumer:
    def __init__(self, name, group_id, 
                bootstrap_servers=DEFAULT_BOOTSTRAP_SERVERS, 
                auto_offset_reset=DEFAULT_AUTO_OFFSET_RESET,
                enable_auto_commit=DEFAULT_COMMIT_FLAG,
                value_deserializer=DEFAULT_CONSUMER_VALUE_SERIALIZER):
        
        self._consumer = KafkaConsumer(
            name,
            bootstrap_servers=bootstrap_servers,
            auto_offset_reset=auto_offset_reset,
            enable_auto_commit=enable_auto_commit,
            group_id=group_id,
            value_deserializer=value_deserializer)