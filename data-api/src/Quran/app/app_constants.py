from json import (dumps, loads)

DEFAULT_BOOTSTRAP_SERVERS = ['localhost:9092']
DEFAULT_AUTO_OFFSET_RESET = 'earliest'
DEFAULT_COMMIT_FLAG = True
DEFAULT_PRODUCER_VALUE_SERIALIZER = lambda x: dumps(x).encode('utf-8')
DEFAULT_CONSUMER_VALUE_SERIALIZER = lambda x: loads(x.decode('utf-8'))
