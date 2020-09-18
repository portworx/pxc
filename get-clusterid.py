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
    # Cluster connection
    clusters = api_pb2_grpc.OpenStorageClusterStub(channel)
    ic_resp = clusters.InspectCurrent(api_pb2.SdkClusterInspectCurrentRequest())
    print(ic_resp.cluster.id)

except grpc.RpcError as e:
    print('Failed: code={0} msg={1}'.format(e.code(), e.details()))

