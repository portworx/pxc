#
# Install:
#   pip3 install --user --upgrade libopenstorage-openstorage
#
import grpc

from openstorage import api_pb2
from openstorage import api_pb2_grpc
from openstorage import connector

c = connector.Connector()
channel = c.connect()

try:
    # Create a volume
    volumes = api_pb2_grpc.OpenStorageVolumeStub(channel)
    ids = volumes.Enumerate(api_pb2.SdkVolumeEnumerateRequest())
    print(ids.volume_ids)
    for id in ids.volume_ids:
        inspect = volumes.Inspect(api_pb2.SdkVolumeInspectRequest(volume_id=id))
        print('name={0} usage={1}'.format(inspect.name, inspect.volume.spec.size))
except grpc.RpcError as e:
    print('Failed: code={0} msg={1}'.format(e.code(), e.details()))

