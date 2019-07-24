/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    httd://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package portworx

var dummyInputJson = `{
  "Vols": [
    {
      "volume": {
        "id": "890671739599487688",
        "source": {},
        "locator": {
          "name": "tp3"
        },
        "ctime": {
          "seconds": 1564053601,
          "nanos": 966147737
        },
        "spec": {
          "size": 5368709120,
          "format": 2,
          "block_size": 4096,
          "ha_level": 1,
          "cos": 1,
          "replica_set": {},
          "aggregation_level": 3,
          "scale": 1,
          "queue_depth": 128,
          "force_unsupported_fs_type": true,
          "io_strategy": {}
        },
        "usage": 1425408,
        "last_scan": {
          "seconds": 1564053601,
          "nanos": 966147737
        },
        "format": 2,
        "status": 2,
        "state": 3,
        "attached_on": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
        "device_path": "/dev/pxd/pxd890671739599487688",
        "replica_sets": [
          {
            "nodes": [
              "4e056193-1f08-44f3-ac69-edf194937edd"
            ]
          },
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693"
            ]
          },
          {
            "nodes": [
              "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "FullResyncBlocks": "[{2 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReadQuorum": "1",
              "ReadSet": "[2]",
              "ReplNodePools": "0",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "4e056193-1f08-44f3-ac69-edf194937edd",
              "ReplicaSetCurr": "[2]",
              "ReplicaSetCurrMid": "4e056193-1f08-44f3-ac69-edf194937edd",
              "ReplicaSetNext": "[2]",
              "ReplicaSetNextMid": "4e056193-1f08-44f3-ac69-edf194937edd",
              "ResyncBlocks": "[{2 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "RuntimeState": "clean",
              "TimestampBlocksPerNode": "[0 0 0 0 0]",
              "TimestampBlocksTotal": "0",
              "WriteQuorum": "1",
              "WriteSet": "[2]"
            }
          },
          {
            "runtime_state": {
              "FullResyncBlocks": "[{0 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "ID": "1",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReadQuorum": "1",
              "ReadSet": "[0]",
              "ReplNodePools": "1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ReplicaSetCurr": "[0]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ReplicaSetNext": "[0]",
              "ReplicaSetNextMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ResyncBlocks": "[{0 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "RuntimeState": "clean",
              "TimestampBlocksPerNode": "[0 0 0 0 0]",
              "TimestampBlocksTotal": "0",
              "WriteQuorum": "1",
              "WriteSet": "[0]"
            }
          },
          {
            "runtime_state": {
              "FullResyncBlocks": "[{1 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "ID": "2",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReadQuorum": "1",
              "ReadSet": "[1]",
              "ReplNodePools": "0",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetCurr": "[1]",
              "ReplicaSetCurrMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetNext": "[1]",
              "ReplicaSetNextMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ResyncBlocks": "[{1 0} {-1 0} {-1 0} {-1 0} {-1 0}]",
              "RuntimeState": "clean",
              "TimestampBlocksPerNode": "[0 0 0 0 0]",
              "TimestampBlocksTotal": "0",
              "WriteQuorum": "1",
              "WriteSet": "[1]"
            }
          }
        ]
      },
      "name": "tp3"
    },
    {
      "volume": {
        "id": "191224272263020605",
        "source": {},
        "locator": {
          "name": "tp1"
        },
        "ctime": {
          "seconds": 1564053287,
          "nanos": 474602314
        },
        "spec": {
          "size": 5368709120,
          "format": 2,
          "block_size": 4096,
          "ha_level": 1,
          "cos": 1,
          "replica_set": {},
          "aggregation_level": 1,
          "snapshot_schedule": "- freq: periodic\n  period: 1200000000000\n  retain: 7\n- freq: daily\n  hour: 10\n  retain: 10\n- freq: weekly\n  weekday: 1\n  hour: 1\n  minute: 25\n  retain: 3\n- freq: monthly\n  day: 1\n  minute: 30\n  retain: 3\n",
          "scale": 1,
          "queue_depth": 128,
          "force_unsupported_fs_type": true,
          "io_strategy": {}
        },
        "last_scan": {
          "seconds": 1564053287,
          "nanos": 474602314
        },
        "format": 2,
        "status": 2,
        "state": 4,
        "attached_state": 2,
        "replica_sets": [
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReplNodePools": "0",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ReplicaSetCurr": "[0]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "RuntimeState": "clean"
            }
          }
        ]
      },
      "name": "tp1"
    },
    {
      "volume": {
        "id": "653648362965696033",
        "source": {},
        "locator": {
          "name": "tp2"
        },
        "ctime": {
          "seconds": 1564053554,
          "nanos": 967711758
        },
        "spec": {
          "size": 5368709120,
          "format": 2,
          "block_size": 4096,
          "ha_level": 1,
          "cos": 1,
          "replica_set": {},
          "aggregation_level": 2,
          "scale": 1,
          "queue_depth": 128,
          "force_unsupported_fs_type": true,
          "io_strategy": {}
        },
        "last_scan": {
          "seconds": 1564053554,
          "nanos": 967711758
        },
        "format": 2,
        "status": 2,
        "state": 3,
        "attached_on": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
        "attached_state": 1,
        "device_path": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
        "replica_sets": [
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693"
            ]
          },
          {
            "nodes": [
              "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReplNodePools": "1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ReplicaSetCurr": "[0]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "RuntimeState": "clean"
            }
          },
          {
            "runtime_state": {
              "ID": "1",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReplNodePools": "1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetCurr": "[1]",
              "ReplicaSetCurrMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "RuntimeState": "clean"
            }
          }
        ]
      },
      "name": "tp2"
    },
    {
      "volume": {
        "id": "5386813725974049",
        "source": {},
        "locator": {
          "name": "pvc-6fc1fe2d-25f4-40b0-a616-04c019572154",
          "volume_labels": {
            "namespace": "wp1",
            "priority_io": "high",
            "pvc": "mysql-pvc-1",
            "repl": "2"
          }
        },
        "ctime": {
          "seconds": 1564044598,
          "nanos": 724669964
        },
        "spec": {
          "size": 21474836480,
          "format": 2,
          "block_size": 4096,
          "ha_level": 2,
          "cos": 3,
          "volume_labels": {
            "namespace": "wp1",
            "priority_io": "high",
            "pvc": "mysql-pvc-1",
            "repl": "2"
          },
          "aggregation_level": 1,
          "queue_depth": 128
        },
        "usage": 33693696,
        "last_scan": {
          "seconds": 1564044598,
          "nanos": 724669964
        },
        "format": 2,
        "status": 2,
        "state": 3,
        "attached_on": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
        "device_path": "/dev/pxd/pxd5386813725974049",
        "attach_path": [
          "/var/lib/kubelet/pods/d37b56a7-3b32-46d0-9e95-849b45bf7c98/volumes/kubernetes.io~portworx-volume/pvc-6fc1fe2d-25f4-40b0-a616-04c019572154"
        ],
        "replica_sets": [
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "FullResyncBlocks": "[{0 0} {1 0} {-1 0} {-1 0} {-1 0}]",
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReadQuorum": "1",
              "ReadSet": "[0 1]",
              "ReplNodePools": "1,1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetCurr": "[0 1]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetNext": "[0 1]",
              "ReplicaSetNextMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ResyncBlocks": "[{0 0} {1 155} {-1 0} {-1 0} {-1 0}]",
              "RuntimeState": "clean",
              "TimestampBlocksPerNode": "[0 0 0 0 0]",
              "TimestampBlocksTotal": "0",
              "WriteQuorum": "2",
              "WriteSet": "[0 1]"
            }
          }
        ]
      },
      "name": "pvc-6fc1fe2d-25f4-40b0-a616-04c019572154",
      "labels": {
        "namespace": "wp1",
        "priority_io": "high",
        "pvc": "mysql-pvc-1",
        "repl": "2"
      }
    },
    {
      "volume": {
        "id": "1122787946485635103",
        "source": {
          "parent": "653648362965696033"
        },
        "readonly": true,
        "locator": {
          "name": "tp2-snap"
        },
        "ctime": {
          "seconds": 1564209302,
          "nanos": 45463894
        },
        "spec": {
          "size": 5368709120,
          "format": 2,
          "block_size": 4096,
          "ha_level": 1,
          "cos": 1,
          "replica_set": {},
          "aggregation_level": 2,
          "scale": 1,
          "sticky": true,
          "queue_depth": 128,
          "force_unsupported_fs_type": true,
          "io_strategy": {}
        },
        "last_scan": {
          "seconds": 1564209301,
          "nanos": 991017473
        },
        "format": 2,
        "status": 2,
        "state": 4,
        "replica_sets": [
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693"
            ]
          },
          {
            "nodes": [
              "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReplNodePools": "1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "ReplicaSetCurr": "[0]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "RuntimeState": "clean"
            }
          },
          {
            "runtime_state": {
              "ID": "1",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReplNodePools": "1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "ReplicaSetCurr": "[1]",
              "ReplicaSetCurrMid": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
              "RuntimeState": "clean"
            }
          }
        ]
      },
      "name": "tp2-snap"
    },
    {
      "volume": {
        "id": "971700916980434894",
        "source": {},
        "locator": {
          "name": "pvc-34d0f15c-65b9-4229-8b3e-b7bb912e382f",
          "volume_labels": {
            "namespace": "wp1",
            "priority_io": "high",
            "pvc": "wp-pv-claim",
            "repl": "2",
            "shared": "true"
          }
        },
        "ctime": {
          "seconds": 1564044599,
          "nanos": 56622403
        },
        "spec": {
          "size": 10737418240,
          "format": 2,
          "block_size": 4096,
          "ha_level": 2,
          "cos": 3,
          "volume_labels": {
            "app": "wordpress",
            "namespace": "wp1",
            "priority_io": "high",
            "pvc": "wp-pv-claim",
            "repl": "2",
            "shared": "true"
          },
          "shared": true,
          "aggregation_level": 1,
          "queue_depth": 128
        },
        "usage": 50155520,
        "last_scan": {
          "seconds": 1564044599,
          "nanos": 56622403
        },
        "format": 2,
        "status": 2,
        "state": 3,
        "attached_on": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
        "device_path": "/dev/pxd/pxd971700916980434894",
        "attach_path": [
          "70.0.87.233",
          "70.0.87.200"
        ],
        "replica_sets": [
          {
            "nodes": [
              "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
              "4e056193-1f08-44f3-ac69-edf194937edd"
            ]
          }
        ],
        "runtime_state": [
          {
            "runtime_state": {
              "FullResyncBlocks": "[{0 0} {2 0} {-1 0} {-1 0} {-1 0}]",
              "ID": "0",
              "PXReplReAddNodeMid": "",
              "PXReplReAddPools": "",
              "ReadQuorum": "1",
              "ReadSet": "[0 2]",
              "ReplNodePools": "1,1",
              "ReplRemoveMids": "",
              "ReplicaSetCreateMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,4e056193-1f08-44f3-ac69-edf194937edd",
              "ReplicaSetCurr": "[0 2]",
              "ReplicaSetCurrMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,4e056193-1f08-44f3-ac69-edf194937edd",
              "ReplicaSetNext": "[0 2]",
              "ReplicaSetNextMid": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693,4e056193-1f08-44f3-ac69-edf194937edd",
              "ResyncBlocks": "[{0 16} {2 0} {-1 0} {-1 0} {-1 0}]",
              "RuntimeState": "clean",
              "TimestampBlocksPerNode": "[0 0 0 0 0]",
              "TimestampBlocksTotal": "0",
              "WriteQuorum": "2",
              "WriteSet": "[0 2]"
            }
          }
        ]
      },
      "name": "pvc-34d0f15c-65b9-4229-8b3e-b7bb912e382f",
      "labels": {
        "namespace": "wp1",
        "priority_io": "high",
        "pvc": "wp-pv-claim",
        "repl": "2",
        "shared": "true"
      }
    }
  ],
  "Pods": [
    {
      "metadata": {
        "name": "kubernetes-bootcamp-5b48cfdcbd-6r85b",
        "generateName": "kubernetes-bootcamp-5b48cfdcbd-",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/pods/kubernetes-bootcamp-5b48cfdcbd-6r85b",
        "uid": "ff86783c-c88e-4b0a-998d-23b533581936",
        "resourceVersion": "137677",
        "creationTimestamp": "2019-07-10T09:56:33Z",
        "labels": {
          "pod-template-hash": "5b48cfdcbd",
          "run": "kubernetes-bootcamp"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "kubernetes-bootcamp-5b48cfdcbd",
            "uid": "f3299cc3-01ea-4a59-90b2-5067ee006d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-4l828",
            "secret": {
              "secretName": "default-token-4l828",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kubernetes-bootcamp",
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "ports": [
              {
                "containerPort": 8080,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "default-token-4l828",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:34Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:34Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.130",
        "startTime": "2019-07-10T09:56:33Z",
        "containerStatuses": [
          {
            "name": "kubernetes-bootcamp",
            "state": {
              "running": {
                "startedAt": "2019-07-10T09:56:34Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "imageID": "docker-pullable://gcr.io/google-samples/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af",
            "containerID": "docker://6619305acf9da962879b72e7dbd3b671422aec6f4d96b1bccfa8edf94891562f"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kubernetes-bootcamp-5b48cfdcbd-7cksb",
        "generateName": "kubernetes-bootcamp-5b48cfdcbd-",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/pods/kubernetes-bootcamp-5b48cfdcbd-7cksb",
        "uid": "d767e2c6-fbe6-431d-bfae-949ea339a653",
        "resourceVersion": "129645",
        "creationTimestamp": "2019-07-09T10:50:53Z",
        "labels": {
          "app": "v1",
          "pod-template-hash": "5b48cfdcbd",
          "run": "kubernetes-bootcamp"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "kubernetes-bootcamp-5b48cfdcbd",
            "uid": "f3299cc3-01ea-4a59-90b2-5067ee006d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-4l828",
            "secret": {
              "secretName": "default-token-4l828",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kubernetes-bootcamp",
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "ports": [
              {
                "containerPort": 8080,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "default-token-4l828",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:50:53Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:51:08Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:51:08Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:50:53Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.129",
        "startTime": "2019-07-09T10:50:53Z",
        "containerStatuses": [
          {
            "name": "kubernetes-bootcamp",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:51:08Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "imageID": "docker-pullable://gcr.io/google-samples/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af",
            "containerID": "docker://1e227cf4d86b5b92fcb556f996ebfff36afee471a65ae35f127f4becfcbd8d8c"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kubernetes-bootcamp-5b48cfdcbd-bwdw9",
        "generateName": "kubernetes-bootcamp-5b48cfdcbd-",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/pods/kubernetes-bootcamp-5b48cfdcbd-bwdw9",
        "uid": "0d7500bc-445f-407d-8e89-f052373493f4",
        "resourceVersion": "137723",
        "creationTimestamp": "2019-07-10T09:56:33Z",
        "labels": {
          "pod-template-hash": "5b48cfdcbd",
          "run": "kubernetes-bootcamp"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "kubernetes-bootcamp-5b48cfdcbd",
            "uid": "f3299cc3-01ea-4a59-90b2-5067ee006d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-4l828",
            "secret": {
              "secretName": "default-token-4l828",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kubernetes-bootcamp",
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "ports": [
              {
                "containerPort": 8080,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "default-token-4l828",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:55Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:55Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.193",
        "startTime": "2019-07-10T09:56:33Z",
        "containerStatuses": [
          {
            "name": "kubernetes-bootcamp",
            "state": {
              "running": {
                "startedAt": "2019-07-10T09:56:54Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "imageID": "docker-pullable://gcr.io/google-samples/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af",
            "containerID": "docker://f5c16d188f4e361ff21e47439e644d32442e163e3940698bb70025562efe0906"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kubernetes-bootcamp-5b48cfdcbd-nlgsv",
        "generateName": "kubernetes-bootcamp-5b48cfdcbd-",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/pods/kubernetes-bootcamp-5b48cfdcbd-nlgsv",
        "uid": "088de0c6-1ab3-440c-9203-d577b6db1180",
        "resourceVersion": "137712",
        "creationTimestamp": "2019-07-10T09:56:33Z",
        "labels": {
          "pod-template-hash": "5b48cfdcbd",
          "run": "kubernetes-bootcamp"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "kubernetes-bootcamp-5b48cfdcbd",
            "uid": "f3299cc3-01ea-4a59-90b2-5067ee006d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-4l828",
            "secret": {
              "secretName": "default-token-4l828",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kubernetes-bootcamp",
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "ports": [
              {
                "containerPort": 8080,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "default-token-4l828",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:52Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:52Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-10T09:56:33Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "192.168.220.193",
        "startTime": "2019-07-10T09:56:33Z",
        "containerStatuses": [
          {
            "name": "kubernetes-bootcamp",
            "state": {
              "running": {
                "startedAt": "2019-07-10T09:56:51Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google-samples/kubernetes-bootcamp:v1",
            "imageID": "docker-pullable://gcr.io/google-samples/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af",
            "containerID": "docker://a940e800927fe5f540998122cf4881b875309ec8e67714f0d07b7b08f2aeb323"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "alertmanager-portworx-0",
        "generateName": "alertmanager-portworx-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/alertmanager-portworx-0",
        "uid": "297cd775-03a3-4a29-a139-09a1ed356659",
        "resourceVersion": "2161173",
        "creationTimestamp": "2019-07-25T08:38:49Z",
        "labels": {
          "alertmanager": "portworx",
          "app": "alertmanager",
          "controller-revision-hash": "alertmanager-portworx-79c87569d6",
          "statefulset.kubernetes.io/pod-name": "alertmanager-portworx-0"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "StatefulSet",
            "name": "alertmanager-portworx",
            "uid": "34796ccc-5697-4ebe-98fc-388e37286a50",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config-volume",
            "secret": {
              "secretName": "alertmanager-portworx",
              "defaultMode": 420
            }
          },
          {
            "name": "alertmanager-portworx-db",
            "emptyDir": {}
          },
          {
            "name": "default-token-g229d",
            "secret": {
              "secretName": "default-token-g229d",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "alertmanager",
            "image": "quay.io/prometheus/alertmanager:v0.16.1",
            "args": [
              "--config.file=/etc/alertmanager/config/alertmanager.yaml",
              "--cluster.listen-address=[$(POD_IP)]:6783",
              "--storage.path=/alertmanager",
              "--data.retention=120h",
              "--web.listen-address=:9093",
              "--web.route-prefix=/",
              "--cluster.peer=alertmanager-portworx-0.alertmanager-operated.kube-system.svc:6783",
              "--cluster.peer=alertmanager-portworx-1.alertmanager-operated.kube-system.svc:6783",
              "--cluster.peer=alertmanager-portworx-2.alertmanager-operated.kube-system.svc:6783"
            ],
            "ports": [
              {
                "name": "web",
                "containerPort": 9093,
                "protocol": "TCP"
              },
              {
                "name": "mesh",
                "containerPort": 6783,
                "protocol": "TCP"
              }
            ],
            "env": [
              {
                "name": "POD_IP",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "status.podIP"
                  }
                }
              }
            ],
            "resources": {
              "requests": {
                "memory": "200Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config-volume",
                "mountPath": "/etc/alertmanager/config"
              },
              {
                "name": "alertmanager-portworx-db",
                "mountPath": "/alertmanager"
              },
              {
                "name": "default-token-g229d",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/api/v1/status",
                "port": "web",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 3,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 10
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/api/v1/status",
                "port": "web",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 3,
              "timeoutSeconds": 3,
              "periodSeconds": 5,
              "successThreshold": 1,
              "failureThreshold": 10
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "config-reloader",
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "args": [
              "-webhook-url=http://localhost:9093/-/reload",
              "-volume-dir=/etc/alertmanager/config"
            ],
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "25Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "25Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config-volume",
                "readOnly": true,
                "mountPath": "/etc/alertmanager/config"
              },
              {
                "name": "default-token-g229d",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 0,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "hostname": "alertmanager-portworx-0",
        "subdomain": "alertmanager-operated",
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Pending",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z"
          },
          {
            "type": "Ready",
            "status": "False",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z",
            "reason": "ContainersNotReady",
            "message": "containers with unready status: [alertmanager config-reloader]"
          },
          {
            "type": "ContainersReady",
            "status": "False",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z",
            "reason": "ContainersNotReady",
            "message": "containers with unready status: [alertmanager config-reloader]"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:49Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "startTime": "2019-07-25T08:38:59Z",
        "containerStatuses": [
          {
            "name": "alertmanager",
            "state": {
              "waiting": {
                "reason": "ContainerCreating"
              }
            },
            "lastState": {},
            "ready": false,
            "restartCount": 0,
            "image": "quay.io/prometheus/alertmanager:v0.16.1",
            "imageID": ""
          },
          {
            "name": "config-reloader",
            "state": {
              "waiting": {
                "reason": "ContainerCreating"
              }
            },
            "lastState": {},
            "ready": false,
            "restartCount": 0,
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "imageID": ""
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "calico-etcd-n7nrd",
        "generateName": "calico-etcd-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-etcd-n7nrd",
        "uid": "981b8836-89e8-4c9c-8d09-6055be2ba94e",
        "resourceVersion": "1136",
        "creationTimestamp": "2019-07-09T09:41:12Z",
        "labels": {
          "controller-revision-hash": "6b6d974699",
          "k8s-app": "calico-etcd",
          "pod-template-generation": "1"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "calico-etcd",
            "uid": "25bf24df-c44e-49f9-9c42-23497ed58035",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "var-etcd",
            "hostPath": {
              "path": "/var/etcd",
              "type": ""
            }
          },
          {
            "name": "default-token-g229d",
            "secret": {
              "secretName": "default-token-g229d",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-etcd",
            "image": "quay.io/coreos/etcd:v3.1.10",
            "command": [
              "/usr/local/bin/etcd"
            ],
            "args": [
              "--name=calico",
              "--data-dir=/var/etcd/calico-data",
              "--advertise-client-urls=http://$CALICO_ETCD_IP:6666",
              "--listen-client-urls=http://0.0.0.0:6666",
              "--listen-peer-urls=http://0.0.0.0:6667",
              "--auto-compaction-retention=1"
            ],
            "env": [
              {
                "name": "CALICO_ETCD_IP",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "status.podIP"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "var-etcd",
                "mountPath": "/var/etcd"
              },
              {
                "name": "default-token-g229d",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "node-role.kubernetes.io/master": ""
        },
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-199.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.cloudprovider.kubernetes.io/uninitialized",
            "value": "true",
            "effect": "NoSchedule"
          },
          {
            "key": "node-role.kubernetes.io/master",
            "effect": "NoSchedule"
          },
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:12Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:24Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:24Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:12Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:41:12Z",
        "containerStatuses": [
          {
            "name": "calico-etcd",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:23Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/etcd:v3.1.10",
            "imageID": "docker-pullable://quay.io/coreos/etcd@sha256:63b041ee0e2453e8591603a8b19decefc91e58912d1d8cfb2165b2ac24a43dd7",
            "containerID": "docker://a02e970c2e677b41a099942f6ff48730610062266106d5dc1fb1729ec62653a6"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "calico-kube-controllers-6b6f4f7c64-pz5bh",
        "generateName": "calico-kube-controllers-6b6f4f7c64-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-kube-controllers-6b6f4f7c64-pz5bh",
        "uid": "b56caf74-b674-4b6b-a54b-5c6fa88310b8",
        "resourceVersion": "1164",
        "creationTimestamp": "2019-07-09T09:40:50Z",
        "labels": {
          "k8s-app": "calico-kube-controllers",
          "pod-template-hash": "6b6f4f7c64"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "calico-kube-controllers-6b6f4f7c64",
            "uid": "cc0a6efd-0520-4540-8650-d5f8af57f6c2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "calico-kube-controllers-token-6nf9x",
            "secret": {
              "secretName": "calico-kube-controllers-token-6nf9x",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-kube-controllers",
            "image": "quay.io/calico/kube-controllers:v3.1.6",
            "env": [
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "ENABLED_CONTROLLERS",
                "value": "policy,profile,workloadendpoint,node"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "calico-kube-controllers-token-6nf9x",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "calico-kube-controllers",
        "serviceAccount": "calico-kube-controllers",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "key": "node-role.kubernetes.io/master",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:12Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:28Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:28Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:12Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:41:12Z",
        "containerStatuses": [
          {
            "name": "calico-kube-controllers",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:28Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/kube-controllers:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/kube-controllers@sha256:b126c0fc99d8df7319a8c441b6c6fb4525967a41a61086ebb1c7f2cd817c03ae",
            "containerID": "docker://e53d428fca619c2ae3ec480261b3dab8576f51d270bfd0455ded2c9d4dc7b6b4"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "calico-node-6p66p",
        "generateName": "calico-node-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-node-6p66p",
        "uid": "d63b6b3a-6ebf-4abd-9266-188dc4589c7e",
        "resourceVersion": "3108",
        "creationTimestamp": "2019-07-09T10:05:38Z",
        "labels": {
          "controller-revision-hash": "7f9fd56fb",
          "k8s-app": "calico-node",
          "pod-template-generation": "1"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "calico-node",
            "uid": "196e5c3d-f241-4e19-a395-0ac641b0145e",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "var-run-calico",
            "hostPath": {
              "path": "/var/run/calico",
              "type": ""
            }
          },
          {
            "name": "var-lib-calico",
            "hostPath": {
              "path": "/var/lib/calico",
              "type": ""
            }
          },
          {
            "name": "cni-bin-dir",
            "hostPath": {
              "path": "/opt/cni/bin",
              "type": ""
            }
          },
          {
            "name": "cni-net-dir",
            "hostPath": {
              "path": "/etc/cni/net.d",
              "type": ""
            }
          },
          {
            "name": "calico-cni-plugin-token-wdhz4",
            "secret": {
              "secretName": "calico-cni-plugin-token-wdhz4",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-node",
            "image": "quay.io/calico/node:v3.1.6",
            "env": [
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CALICO_NETWORKING_BACKEND",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "calico_backend"
                  }
                }
              },
              {
                "name": "CLUSTER_TYPE",
                "value": "kubeadm,bgp"
              },
              {
                "name": "CALICO_DISABLE_FILE_LOGGING",
                "value": "true"
              },
              {
                "name": "CALICO_K8S_NODE_REF",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              },
              {
                "name": "FELIX_DEFAULTENDPOINTTOHOSTACTION",
                "value": "ACCEPT"
              },
              {
                "name": "CALICO_IPV4POOL_CIDR",
                "value": "192.168.0.0/16"
              },
              {
                "name": "CALICO_IPV4POOL_IPIP",
                "value": "Always"
              },
              {
                "name": "FELIX_IPV6SUPPORT",
                "value": "false"
              },
              {
                "name": "FELIX_IPINIPMTU",
                "value": "1440"
              },
              {
                "name": "FELIX_LOGSEVERITYSCREEN",
                "value": "info"
              },
              {
                "name": "IP",
                "value": "autodetect"
              },
              {
                "name": "FELIX_HEALTHENABLED",
                "value": "true"
              }
            ],
            "resources": {
              "requests": {
                "cpu": "250m"
              }
            },
            "volumeMounts": [
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "var-run-calico",
                "mountPath": "/var/run/calico"
              },
              {
                "name": "var-lib-calico",
                "mountPath": "/var/lib/calico"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/liveness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/readiness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          },
          {
            "name": "install-cni",
            "image": "quay.io/calico/cni:v3.1.6",
            "command": [
              "/install-cni.sh"
            ],
            "env": [
              {
                "name": "CNI_CONF_NAME",
                "value": "10-calico.conflist"
              },
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CNI_NETWORK_CONFIG",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "cni_network_config"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "cni-bin-dir",
                "mountPath": "/host/opt/cni/bin"
              },
              {
                "name": "cni-net-dir",
                "mountPath": "/host/etc/cni/net.d"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 0,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "calico-cni-plugin",
        "serviceAccount": "calico-cni-plugin",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-200.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:05:38Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:32Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:32Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:05:38Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "70.0.87.200",
        "startTime": "2019-07-09T10:05:38Z",
        "containerStatuses": [
          {
            "name": "calico-node",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:06:20Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/node:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/node@sha256:dd4bd919e785084b84d5e4f89d7bc0125b3b021f2e83ea7fe23c9662789fe6df",
            "containerID": "docker://98da2581270ccdf554298244370be8a7837fde0ad3e33cd288185ede38578186"
          },
          {
            "name": "install-cni",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:06:24Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/cni:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/cni@sha256:f9d231b82dde6ce74fd7b7dfbbc84f9a0b23e46e606f2b600ff3fed8ed940b4f",
            "containerID": "docker://39da94a7a6e27122bbc0022a35fdf57f8e31a3b70008d311860f27818f169401"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "calico-node-7kgjs",
        "generateName": "calico-node-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-node-7kgjs",
        "uid": "f4d15142-2543-4556-a133-447eb0f48843",
        "resourceVersion": "3295",
        "creationTimestamp": "2019-07-09T10:06:53Z",
        "labels": {
          "controller-revision-hash": "7f9fd56fb",
          "k8s-app": "calico-node",
          "pod-template-generation": "1"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "calico-node",
            "uid": "196e5c3d-f241-4e19-a395-0ac641b0145e",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "var-run-calico",
            "hostPath": {
              "path": "/var/run/calico",
              "type": ""
            }
          },
          {
            "name": "var-lib-calico",
            "hostPath": {
              "path": "/var/lib/calico",
              "type": ""
            }
          },
          {
            "name": "cni-bin-dir",
            "hostPath": {
              "path": "/opt/cni/bin",
              "type": ""
            }
          },
          {
            "name": "cni-net-dir",
            "hostPath": {
              "path": "/etc/cni/net.d",
              "type": ""
            }
          },
          {
            "name": "calico-cni-plugin-token-wdhz4",
            "secret": {
              "secretName": "calico-cni-plugin-token-wdhz4",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-node",
            "image": "quay.io/calico/node:v3.1.6",
            "env": [
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CALICO_NETWORKING_BACKEND",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "calico_backend"
                  }
                }
              },
              {
                "name": "CLUSTER_TYPE",
                "value": "kubeadm,bgp"
              },
              {
                "name": "CALICO_DISABLE_FILE_LOGGING",
                "value": "true"
              },
              {
                "name": "CALICO_K8S_NODE_REF",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              },
              {
                "name": "FELIX_DEFAULTENDPOINTTOHOSTACTION",
                "value": "ACCEPT"
              },
              {
                "name": "CALICO_IPV4POOL_CIDR",
                "value": "192.168.0.0/16"
              },
              {
                "name": "CALICO_IPV4POOL_IPIP",
                "value": "Always"
              },
              {
                "name": "FELIX_IPV6SUPPORT",
                "value": "false"
              },
              {
                "name": "FELIX_IPINIPMTU",
                "value": "1440"
              },
              {
                "name": "FELIX_LOGSEVERITYSCREEN",
                "value": "info"
              },
              {
                "name": "IP",
                "value": "autodetect"
              },
              {
                "name": "FELIX_HEALTHENABLED",
                "value": "true"
              }
            ],
            "resources": {
              "requests": {
                "cpu": "250m"
              }
            },
            "volumeMounts": [
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "var-run-calico",
                "mountPath": "/var/run/calico"
              },
              {
                "name": "var-lib-calico",
                "mountPath": "/var/lib/calico"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/liveness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/readiness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          },
          {
            "name": "install-cni",
            "image": "quay.io/calico/cni:v3.1.6",
            "command": [
              "/install-cni.sh"
            ],
            "env": [
              {
                "name": "CNI_CONF_NAME",
                "value": "10-calico.conflist"
              },
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CNI_NETWORK_CONFIG",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "cni_network_config"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "cni-bin-dir",
                "mountPath": "/host/opt/cni/bin"
              },
              {
                "name": "cni-net-dir",
                "mountPath": "/host/etc/cni/net.d"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 0,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "calico-cni-plugin",
        "serviceAccount": "calico-cni-plugin",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-203.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:53Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:39Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:39Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:53Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "70.0.87.203",
        "startTime": "2019-07-09T10:06:53Z",
        "containerStatuses": [
          {
            "name": "calico-node",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:07:28Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/node:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/node@sha256:dd4bd919e785084b84d5e4f89d7bc0125b3b021f2e83ea7fe23c9662789fe6df",
            "containerID": "docker://a2cdd876faf8fcf845fc2deac2ba6f952de5d064e3c9911200d08e2e7ac584ab"
          },
          {
            "name": "install-cni",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:07:31Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/cni:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/cni@sha256:f9d231b82dde6ce74fd7b7dfbbc84f9a0b23e46e606f2b600ff3fed8ed940b4f",
            "containerID": "docker://3f1822975372bd489e4c61916deaedd5c3ece3bdff51d5934508a7b2c703a184"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "calico-node-g5hsl",
        "generateName": "calico-node-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-node-g5hsl",
        "uid": "22702dcf-50b0-463b-86db-fdc5f03b16f2",
        "resourceVersion": "3227",
        "creationTimestamp": "2019-07-09T10:06:20Z",
        "labels": {
          "controller-revision-hash": "7f9fd56fb",
          "k8s-app": "calico-node",
          "pod-template-generation": "1"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "calico-node",
            "uid": "196e5c3d-f241-4e19-a395-0ac641b0145e",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "var-run-calico",
            "hostPath": {
              "path": "/var/run/calico",
              "type": ""
            }
          },
          {
            "name": "var-lib-calico",
            "hostPath": {
              "path": "/var/lib/calico",
              "type": ""
            }
          },
          {
            "name": "cni-bin-dir",
            "hostPath": {
              "path": "/opt/cni/bin",
              "type": ""
            }
          },
          {
            "name": "cni-net-dir",
            "hostPath": {
              "path": "/etc/cni/net.d",
              "type": ""
            }
          },
          {
            "name": "calico-cni-plugin-token-wdhz4",
            "secret": {
              "secretName": "calico-cni-plugin-token-wdhz4",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-node",
            "image": "quay.io/calico/node:v3.1.6",
            "env": [
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CALICO_NETWORKING_BACKEND",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "calico_backend"
                  }
                }
              },
              {
                "name": "CLUSTER_TYPE",
                "value": "kubeadm,bgp"
              },
              {
                "name": "CALICO_DISABLE_FILE_LOGGING",
                "value": "true"
              },
              {
                "name": "CALICO_K8S_NODE_REF",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              },
              {
                "name": "FELIX_DEFAULTENDPOINTTOHOSTACTION",
                "value": "ACCEPT"
              },
              {
                "name": "CALICO_IPV4POOL_CIDR",
                "value": "192.168.0.0/16"
              },
              {
                "name": "CALICO_IPV4POOL_IPIP",
                "value": "Always"
              },
              {
                "name": "FELIX_IPV6SUPPORT",
                "value": "false"
              },
              {
                "name": "FELIX_IPINIPMTU",
                "value": "1440"
              },
              {
                "name": "FELIX_LOGSEVERITYSCREEN",
                "value": "info"
              },
              {
                "name": "IP",
                "value": "autodetect"
              },
              {
                "name": "FELIX_HEALTHENABLED",
                "value": "true"
              }
            ],
            "resources": {
              "requests": {
                "cpu": "250m"
              }
            },
            "volumeMounts": [
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "var-run-calico",
                "mountPath": "/var/run/calico"
              },
              {
                "name": "var-lib-calico",
                "mountPath": "/var/lib/calico"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/liveness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/readiness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          },
          {
            "name": "install-cni",
            "image": "quay.io/calico/cni:v3.1.6",
            "command": [
              "/install-cni.sh"
            ],
            "env": [
              {
                "name": "CNI_CONF_NAME",
                "value": "10-calico.conflist"
              },
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CNI_NETWORK_CONFIG",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "cni_network_config"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "cni-bin-dir",
                "mountPath": "/host/opt/cni/bin"
              },
              {
                "name": "cni-net-dir",
                "mountPath": "/host/etc/cni/net.d"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 0,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "calico-cni-plugin",
        "serviceAccount": "calico-cni-plugin",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-233.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:20Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:11Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:11Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:20Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "70.0.87.233",
        "startTime": "2019-07-09T10:06:20Z",
        "containerStatuses": [
          {
            "name": "calico-node",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:07:02Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 1,
                "reason": "Error",
                "startedAt": "2019-07-09T10:06:51Z",
                "finishedAt": "2019-07-09T10:07:01Z",
                "containerID": "docker://0d18336ab9fec4f58e1766d1d4c7a36799658bd869acc0ef6e1ac68c2a24f972"
              }
            },
            "ready": true,
            "restartCount": 1,
            "image": "quay.io/calico/node:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/node@sha256:dd4bd919e785084b84d5e4f89d7bc0125b3b021f2e83ea7fe23c9662789fe6df",
            "containerID": "docker://46542d19e5583578fd7394c3774c95022717de2695e582e73cea2254960c0989"
          },
          {
            "name": "install-cni",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:06:59Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/cni:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/cni@sha256:f9d231b82dde6ce74fd7b7dfbbc84f9a0b23e46e606f2b600ff3fed8ed940b4f",
            "containerID": "docker://2bb1a349487f1b97acc07d9e55004d7d43fe179b14f7df2f368219d44de8a4f1"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "calico-node-r9mrs",
        "generateName": "calico-node-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/calico-node-r9mrs",
        "uid": "ac9a21b4-883a-48b1-a7e9-7f46309306c9",
        "resourceVersion": "1184",
        "creationTimestamp": "2019-07-09T09:40:50Z",
        "labels": {
          "controller-revision-hash": "7f9fd56fb",
          "k8s-app": "calico-node",
          "pod-template-generation": "1"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "calico-node",
            "uid": "196e5c3d-f241-4e19-a395-0ac641b0145e",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "var-run-calico",
            "hostPath": {
              "path": "/var/run/calico",
              "type": ""
            }
          },
          {
            "name": "var-lib-calico",
            "hostPath": {
              "path": "/var/lib/calico",
              "type": ""
            }
          },
          {
            "name": "cni-bin-dir",
            "hostPath": {
              "path": "/opt/cni/bin",
              "type": ""
            }
          },
          {
            "name": "cni-net-dir",
            "hostPath": {
              "path": "/etc/cni/net.d",
              "type": ""
            }
          },
          {
            "name": "calico-cni-plugin-token-wdhz4",
            "secret": {
              "secretName": "calico-cni-plugin-token-wdhz4",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "calico-node",
            "image": "quay.io/calico/node:v3.1.6",
            "env": [
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CALICO_NETWORKING_BACKEND",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "calico_backend"
                  }
                }
              },
              {
                "name": "CLUSTER_TYPE",
                "value": "kubeadm,bgp"
              },
              {
                "name": "CALICO_DISABLE_FILE_LOGGING",
                "value": "true"
              },
              {
                "name": "CALICO_K8S_NODE_REF",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              },
              {
                "name": "FELIX_DEFAULTENDPOINTTOHOSTACTION",
                "value": "ACCEPT"
              },
              {
                "name": "CALICO_IPV4POOL_CIDR",
                "value": "192.168.0.0/16"
              },
              {
                "name": "CALICO_IPV4POOL_IPIP",
                "value": "Always"
              },
              {
                "name": "FELIX_IPV6SUPPORT",
                "value": "false"
              },
              {
                "name": "FELIX_IPINIPMTU",
                "value": "1440"
              },
              {
                "name": "FELIX_LOGSEVERITYSCREEN",
                "value": "info"
              },
              {
                "name": "IP",
                "value": "autodetect"
              },
              {
                "name": "FELIX_HEALTHENABLED",
                "value": "true"
              }
            ],
            "resources": {
              "requests": {
                "cpu": "250m"
              }
            },
            "volumeMounts": [
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "var-run-calico",
                "mountPath": "/var/run/calico"
              },
              {
                "name": "var-lib-calico",
                "mountPath": "/var/lib/calico"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/liveness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/readiness",
                "port": 9099,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          },
          {
            "name": "install-cni",
            "image": "quay.io/calico/cni:v3.1.6",
            "command": [
              "/install-cni.sh"
            ],
            "env": [
              {
                "name": "CNI_CONF_NAME",
                "value": "10-calico.conflist"
              },
              {
                "name": "ETCD_ENDPOINTS",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "etcd_endpoints"
                  }
                }
              },
              {
                "name": "CNI_NETWORK_CONFIG",
                "valueFrom": {
                  "configMapKeyRef": {
                    "name": "calico-config",
                    "key": "cni_network_config"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "cni-bin-dir",
                "mountPath": "/host/opt/cni/bin"
              },
              {
                "name": "cni-net-dir",
                "mountPath": "/host/etc/cni/net.d"
              },
              {
                "name": "calico-cni-plugin-token-wdhz4",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 0,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "calico-cni-plugin",
        "serviceAccount": "calico-cni-plugin",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-199.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:40:50Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:31Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:31Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:40:50Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:40:50Z",
        "containerStatuses": [
          {
            "name": "calico-node",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:29Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 1,
                "reason": "Error",
                "startedAt": "2019-07-09T09:41:08Z",
                "finishedAt": "2019-07-09T09:41:18Z",
                "containerID": "docker://a24b15811b9a884171d78ce02630558954a026fa70d4ca21a3a7faec95d629dd"
              }
            },
            "ready": true,
            "restartCount": 2,
            "image": "quay.io/calico/node:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/node@sha256:dd4bd919e785084b84d5e4f89d7bc0125b3b021f2e83ea7fe23c9662789fe6df",
            "containerID": "docker://cec954ff3aa4933af4592c18f1d65aef5df5026921e4dee3047fd50aaae7e5b0"
          },
          {
            "name": "install-cni",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:02Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/calico/cni:v3.1.6",
            "imageID": "docker-pullable://quay.io/calico/cni@sha256:f9d231b82dde6ce74fd7b7dfbbc84f9a0b23e46e606f2b600ff3fed8ed940b4f",
            "containerID": "docker://a95ece2afd67a7dde69a057b2577b5751f1f3423a3c8255d696c84fefaa5d3e5"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "coredns-5c98db65d4-75j9t",
        "generateName": "coredns-5c98db65d4-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/coredns-5c98db65d4-75j9t",
        "uid": "06443a05-6e38-42d9-9687-e2cb21878dd6",
        "resourceVersion": "1198",
        "creationTimestamp": "2019-07-09T09:32:31Z",
        "labels": {
          "k8s-app": "kube-dns",
          "pod-template-hash": "5c98db65d4"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "coredns-5c98db65d4",
            "uid": "26b1ac59-20ba-45fd-bb60-bc510566fe49",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config-volume",
            "configMap": {
              "name": "coredns",
              "items": [
                {
                  "key": "Corefile",
                  "path": "Corefile"
                }
              ],
              "defaultMode": 420
            }
          },
          {
            "name": "coredns-token-c8947",
            "secret": {
              "secretName": "coredns-token-c8947",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "coredns",
            "image": "k8s.gcr.io/coredns:1.3.1",
            "args": [
              "-conf",
              "/etc/coredns/Corefile"
            ],
            "ports": [
              {
                "name": "dns",
                "containerPort": 53,
                "protocol": "UDP"
              },
              {
                "name": "dns-tcp",
                "containerPort": 53,
                "protocol": "TCP"
              },
              {
                "name": "metrics",
                "containerPort": 9153,
                "protocol": "TCP"
              }
            ],
            "resources": {
              "limits": {
                "memory": "170Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "70Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config-volume",
                "readOnly": true,
                "mountPath": "/etc/coredns"
              },
              {
                "name": "coredns-token-c8947",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 8080,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 60,
              "timeoutSeconds": 5,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 5
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 8080,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "capabilities": {
                "add": [
                  "NET_BIND_SERVICE"
                ],
                "drop": [
                  "all"
                ]
              },
              "readOnlyRootFilesystem": true,
              "allowPrivilegeEscalation": false
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "Default",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "coredns",
        "serviceAccount": "coredns",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "key": "node-role.kubernetes.io/master",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:13Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:38Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:38Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:13Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "192.168.172.66",
        "startTime": "2019-07-09T09:41:13Z",
        "containerStatuses": [
          {
            "name": "coredns",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:29Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/coredns:1.3.1",
            "imageID": "docker-pullable://k8s.gcr.io/coredns@sha256:02382353821b12c21b062c59184e227e001079bb13ebd01f9d3270ba0fcbf1e4",
            "containerID": "docker://d9390bf47c28f3566953be01067ccc04fe0d4c2ef77de9cf105430f4489fe86c"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "coredns-5c98db65d4-zsf7g",
        "generateName": "coredns-5c98db65d4-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/coredns-5c98db65d4-zsf7g",
        "uid": "79c69c72-886a-4a09-91e9-0b5622afac2e",
        "resourceVersion": "1194",
        "creationTimestamp": "2019-07-09T09:32:31Z",
        "labels": {
          "k8s-app": "kube-dns",
          "pod-template-hash": "5c98db65d4"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "coredns-5c98db65d4",
            "uid": "26b1ac59-20ba-45fd-bb60-bc510566fe49",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config-volume",
            "configMap": {
              "name": "coredns",
              "items": [
                {
                  "key": "Corefile",
                  "path": "Corefile"
                }
              ],
              "defaultMode": 420
            }
          },
          {
            "name": "coredns-token-c8947",
            "secret": {
              "secretName": "coredns-token-c8947",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "coredns",
            "image": "k8s.gcr.io/coredns:1.3.1",
            "args": [
              "-conf",
              "/etc/coredns/Corefile"
            ],
            "ports": [
              {
                "name": "dns",
                "containerPort": 53,
                "protocol": "UDP"
              },
              {
                "name": "dns-tcp",
                "containerPort": 53,
                "protocol": "TCP"
              },
              {
                "name": "metrics",
                "containerPort": 9153,
                "protocol": "TCP"
              }
            ],
            "resources": {
              "limits": {
                "memory": "170Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "70Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config-volume",
                "readOnly": true,
                "mountPath": "/etc/coredns"
              },
              {
                "name": "coredns-token-c8947",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 8080,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 60,
              "timeoutSeconds": 5,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 5
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 8080,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "capabilities": {
                "add": [
                  "NET_BIND_SERVICE"
                ],
                "drop": [
                  "all"
                ]
              },
              "readOnlyRootFilesystem": true,
              "allowPrivilegeEscalation": false
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "Default",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "coredns",
        "serviceAccount": "coredns",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "key": "node-role.kubernetes.io/master",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:13Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:38Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:38Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:41:13Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "192.168.172.65",
        "startTime": "2019-07-09T09:41:13Z",
        "containerStatuses": [
          {
            "name": "coredns",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:41:29Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/coredns:1.3.1",
            "imageID": "docker-pullable://k8s.gcr.io/coredns@sha256:02382353821b12c21b062c59184e227e001079bb13ebd01f9d3270ba0fcbf1e4",
            "containerID": "docker://ac2c63d5fa59f879b4c2867a36d0c89e28ff8ae310d4fa49d195d8f3d59aa9d6"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "etcd-ip-70-0-87-199.brbnca.spcsdns.net",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/etcd-ip-70-0-87-199.brbnca.spcsdns.net",
        "uid": "c016112f-e6d8-4b11-af59-632ffc98ccf7",
        "resourceVersion": "479",
        "creationTimestamp": "2019-07-09T09:33:45Z",
        "labels": {
          "component": "etcd",
          "tier": "control-plane"
        },
        "annotations": {
          "kubernetes.io/config.hash": "63794fdaa9b5a55d3b256c1b1cf9d5c9",
          "kubernetes.io/config.mirror": "63794fdaa9b5a55d3b256c1b1cf9d5c9",
          "kubernetes.io/config.seen": "2019-07-09T09:31:54.302794297Z",
          "kubernetes.io/config.source": "file"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "etcd-certs",
            "hostPath": {
              "path": "/etc/kubernetes/pki/etcd",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "etcd-data",
            "hostPath": {
              "path": "/var/lib/etcd",
              "type": "DirectoryOrCreate"
            }
          }
        ],
        "containers": [
          {
            "name": "etcd",
            "image": "k8s.gcr.io/etcd:3.3.10",
            "command": [
              "etcd",
              "--advertise-client-urls=https://70.0.87.199:2379",
              "--cert-file=/etc/kubernetes/pki/etcd/server.crt",
              "--client-cert-auth=true",
              "--data-dir=/var/lib/etcd",
              "--initial-advertise-peer-urls=https://70.0.87.199:2380",
              "--initial-cluster=ip-70-0-87-199.brbnca.spcsdns.net=https://70.0.87.199:2380",
              "--key-file=/etc/kubernetes/pki/etcd/server.key",
              "--listen-client-urls=https://127.0.0.1:2379,https://70.0.87.199:2379",
              "--listen-peer-urls=https://70.0.87.199:2380",
              "--name=ip-70-0-87-199.brbnca.spcsdns.net",
              "--peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt",
              "--peer-client-cert-auth=true",
              "--peer-key-file=/etc/kubernetes/pki/etcd/peer.key",
              "--peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt",
              "--snapshot-count=10000",
              "--trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt"
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "etcd-data",
                "mountPath": "/var/lib/etcd"
              },
              {
                "name": "etcd-certs",
                "mountPath": "/etc/kubernetes/pki/etcd"
              }
            ],
            "livenessProbe": {
              "exec": {
                "command": [
                  "/bin/sh",
                  "-ec",
                  "ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/healthcheck-client.crt --key=/etc/kubernetes/pki/etcd/healthcheck-client.key get foo"
                ]
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 8
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoExecute"
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:32:15Z",
        "containerStatuses": [
          {
            "name": "etcd",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:32:16Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/etcd:3.3.10",
            "imageID": "docker-pullable://k8s.gcr.io/etcd@sha256:17da501f5d2a675be46040422a27b7cc21b8a43895ac998b171db1c346f361f7",
            "containerID": "docker://996f3088b96acd5d37579f2640c35daf7f58300d3117dbf9c6b37237be5ba66a"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "grafana-798c6bc5f8-lxhqx",
        "generateName": "grafana-798c6bc5f8-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/grafana-798c6bc5f8-lxhqx",
        "uid": "d1c00419-742f-42b1-b736-5e2e4caf053d",
        "resourceVersion": "2161466",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "app": "grafana",
          "pod-template-hash": "798c6bc5f8"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "grafana-798c6bc5f8",
            "uid": "6aeb68a0-1c0d-44fa-ae04-ebe5f950bdbd",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "grafana-source-config",
            "configMap": {
              "name": "grafana-source-config",
              "defaultMode": 420
            }
          },
          {
            "name": "grafana-dash-config",
            "configMap": {
              "name": "grafana-dashboard-config",
              "defaultMode": 420
            }
          },
          {
            "name": "dashboard-templates",
            "configMap": {
              "name": "grafana-dashboards",
              "defaultMode": 420
            }
          },
          {
            "name": "default-token-g229d",
            "secret": {
              "secretName": "default-token-g229d",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "grafana",
            "image": "grafana/grafana:5.0.4",
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "100Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "100Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "grafana-dash-config",
                "mountPath": "/etc/grafana/provisioning/dashboards"
              },
              {
                "name": "dashboard-templates",
                "mountPath": "/var/lib/grafana/dashboards"
              },
              {
                "name": "grafana-source-config",
                "mountPath": "/etc/grafana/provisioning/datasources"
              },
              {
                "name": "default-token-g229d",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "readinessProbe": {
              "httpGet": {
                "path": "/login",
                "port": 3000,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:01Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:24Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:24Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.199",
        "startTime": "2019-07-25T08:29:01Z",
        "containerStatuses": [
          {
            "name": "grafana",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:39:37Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "grafana/grafana:5.0.4",
            "imageID": "docker-pullable://grafana/grafana@sha256:9c66c7c01a6bf56023126a0b6f933f4966e8ee795c5f76fa2ad81b3c6dadc1c9",
            "containerID": "docker://cf099569dd71d7d57a107c0164c28ce8f8d28c107ce9c82f5f8272f38b21f338"
          }
        ],
        "qosClass": "Guaranteed"
      }
    },
    {
      "metadata": {
        "name": "kube-apiserver-ip-70-0-87-199.brbnca.spcsdns.net",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-apiserver-ip-70-0-87-199.brbnca.spcsdns.net",
        "uid": "da564bb5-7999-4d0e-ad1f-0424b1c6bd4e",
        "resourceVersion": "462",
        "creationTimestamp": "2019-07-09T09:33:41Z",
        "labels": {
          "component": "kube-apiserver",
          "tier": "control-plane"
        },
        "annotations": {
          "kubernetes.io/config.hash": "ddfba46cc24e3a5e700bdf647db2a3c5",
          "kubernetes.io/config.mirror": "ddfba46cc24e3a5e700bdf647db2a3c5",
          "kubernetes.io/config.seen": "2019-07-09T09:31:54.302796341Z",
          "kubernetes.io/config.source": "file"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "ca-certs",
            "hostPath": {
              "path": "/etc/ssl/certs",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "etc-pki",
            "hostPath": {
              "path": "/etc/pki",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "k8s-certs",
            "hostPath": {
              "path": "/etc/kubernetes/pki",
              "type": "DirectoryOrCreate"
            }
          }
        ],
        "containers": [
          {
            "name": "kube-apiserver",
            "image": "k8s.gcr.io/kube-apiserver:v1.15.0",
            "command": [
              "kube-apiserver",
              "--advertise-address=70.0.87.199",
              "--allow-privileged=true",
              "--authorization-mode=Node,RBAC",
              "--client-ca-file=/etc/kubernetes/pki/ca.crt",
              "--enable-admission-plugins=NodeRestriction",
              "--enable-bootstrap-token-auth=true",
              "--etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt",
              "--etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt",
              "--etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key",
              "--etcd-servers=https://127.0.0.1:2379",
              "--insecure-port=0",
              "--kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt",
              "--kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key",
              "--kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname",
              "--proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt",
              "--proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key",
              "--requestheader-allowed-names=front-proxy-client",
              "--requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt",
              "--requestheader-extra-headers-prefix=X-Remote-Extra-",
              "--requestheader-group-headers=X-Remote-Group",
              "--requestheader-username-headers=X-Remote-User",
              "--secure-port=6443",
              "--service-account-key-file=/etc/kubernetes/pki/sa.pub",
              "--service-cluster-ip-range=10.96.0.0/12",
              "--tls-cert-file=/etc/kubernetes/pki/apiserver.crt",
              "--tls-private-key-file=/etc/kubernetes/pki/apiserver.key"
            ],
            "resources": {
              "requests": {
                "cpu": "250m"
              }
            },
            "volumeMounts": [
              {
                "name": "ca-certs",
                "readOnly": true,
                "mountPath": "/etc/ssl/certs"
              },
              {
                "name": "etc-pki",
                "readOnly": true,
                "mountPath": "/etc/pki"
              },
              {
                "name": "k8s-certs",
                "readOnly": true,
                "mountPath": "/etc/kubernetes/pki"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 6443,
                "host": "70.0.87.199",
                "scheme": "HTTPS"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 8
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoExecute"
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:32:15Z",
        "containerStatuses": [
          {
            "name": "kube-apiserver",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:32:16Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-apiserver:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-apiserver@sha256:8484f7128c5b3f0dddcf04bcc9edbfd6b570fde005d874a6240e23583beb3fcb",
            "containerID": "docker://9d8ec767ce251fe4f48675b9d69ec734a1f5c64ede9b3a9d153d0eb35df6ed70"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "kube-controller-manager-ip-70-0-87-199.brbnca.spcsdns.net",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-controller-manager-ip-70-0-87-199.brbnca.spcsdns.net",
        "uid": "df5c5c6f-9578-4b16-a628-773938f06c80",
        "resourceVersion": "463",
        "creationTimestamp": "2019-07-09T09:33:40Z",
        "labels": {
          "component": "kube-controller-manager",
          "tier": "control-plane"
        },
        "annotations": {
          "kubernetes.io/config.hash": "5455f01beae5568afe1b7ea9d7d0012e",
          "kubernetes.io/config.mirror": "5455f01beae5568afe1b7ea9d7d0012e",
          "kubernetes.io/config.seen": "2019-07-09T09:31:54.302798382Z",
          "kubernetes.io/config.source": "file"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "ca-certs",
            "hostPath": {
              "path": "/etc/ssl/certs",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "etc-pki",
            "hostPath": {
              "path": "/etc/pki",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "k8s-certs",
            "hostPath": {
              "path": "/etc/kubernetes/pki",
              "type": "DirectoryOrCreate"
            }
          },
          {
            "name": "kubeconfig",
            "hostPath": {
              "path": "/etc/kubernetes/controller-manager.conf",
              "type": "FileOrCreate"
            }
          }
        ],
        "containers": [
          {
            "name": "kube-controller-manager",
            "image": "k8s.gcr.io/kube-controller-manager:v1.15.0",
            "command": [
              "kube-controller-manager",
              "--authentication-kubeconfig=/etc/kubernetes/controller-manager.conf",
              "--authorization-kubeconfig=/etc/kubernetes/controller-manager.conf",
              "--bind-address=127.0.0.1",
              "--client-ca-file=/etc/kubernetes/pki/ca.crt",
              "--cluster-signing-cert-file=/etc/kubernetes/pki/ca.crt",
              "--cluster-signing-key-file=/etc/kubernetes/pki/ca.key",
              "--controllers=*,bootstrapsigner,tokencleaner",
              "--kubeconfig=/etc/kubernetes/controller-manager.conf",
              "--leader-elect=true",
              "--requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt",
              "--root-ca-file=/etc/kubernetes/pki/ca.crt",
              "--service-account-private-key-file=/etc/kubernetes/pki/sa.key",
              "--use-service-account-credentials=true"
            ],
            "resources": {
              "requests": {
                "cpu": "200m"
              }
            },
            "volumeMounts": [
              {
                "name": "ca-certs",
                "readOnly": true,
                "mountPath": "/etc/ssl/certs"
              },
              {
                "name": "etc-pki",
                "readOnly": true,
                "mountPath": "/etc/pki"
              },
              {
                "name": "k8s-certs",
                "readOnly": true,
                "mountPath": "/etc/kubernetes/pki"
              },
              {
                "name": "kubeconfig",
                "readOnly": true,
                "mountPath": "/etc/kubernetes/controller-manager.conf"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10252,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 8
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoExecute"
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:32:15Z",
        "containerStatuses": [
          {
            "name": "kube-controller-manager",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:32:15Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-controller-manager:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-controller-manager@sha256:835f32a5cdb30e86f35675dd91f9c7df01d48359ab8b51c1df866a2c7ea2e870",
            "containerID": "docker://55a06ed8dc273e9aa3ce9e28636f6b726ce5b92c8453e39904953f898f243361"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "kube-proxy-2m54p",
        "generateName": "kube-proxy-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-proxy-2m54p",
        "uid": "4f4c1af8-7618-4e0e-bead-705ff893dac4",
        "resourceVersion": "3185",
        "creationTimestamp": "2019-07-09T10:06:20Z",
        "labels": {
          "controller-revision-hash": "7bdbc788b8",
          "k8s-app": "kube-proxy",
          "pod-template-generation": "1"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "kube-proxy",
            "uid": "b3c72459-1656-49c5-a068-32494f57f5b2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "kube-proxy",
            "configMap": {
              "name": "kube-proxy",
              "defaultMode": 420
            }
          },
          {
            "name": "xtables-lock",
            "hostPath": {
              "path": "/run/xtables.lock",
              "type": "FileOrCreate"
            }
          },
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "kube-proxy-token-wfz8m",
            "secret": {
              "secretName": "kube-proxy-token-wfz8m",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kube-proxy",
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "command": [
              "/usr/local/bin/kube-proxy",
              "--config=/var/lib/kube-proxy/config.conf",
              "--hostname-override=$(NODE_NAME)"
            ],
            "env": [
              {
                "name": "NODE_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-proxy",
                "mountPath": "/var/lib/kube-proxy"
              },
              {
                "name": "xtables-lock",
                "mountPath": "/run/xtables.lock"
              },
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "kube-proxy-token-wfz8m",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "kube-proxy",
        "serviceAccount": "kube-proxy",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-233.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priorityClassName": "system-node-critical",
        "priority": 2000001000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:20Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:54Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:54Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:20Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "70.0.87.233",
        "startTime": "2019-07-09T10:06:20Z",
        "containerStatuses": [
          {
            "name": "kube-proxy",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:06:54Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-proxy@sha256:4ef8ca8fa1fbe311f3e6c5d6123d19f48a7bc62984ea8acfc8da8fbdac52bff7",
            "containerID": "docker://942bc2ce55f358c934b7f7453170697fcd7fdecdb3747bab744d222e19bbc37f"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kube-proxy-2rfqm",
        "generateName": "kube-proxy-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-proxy-2rfqm",
        "uid": "a8e36e61-dea8-4f64-9868-7bf4d98cf0f6",
        "resourceVersion": "368",
        "creationTimestamp": "2019-07-09T09:32:31Z",
        "labels": {
          "controller-revision-hash": "7bdbc788b8",
          "k8s-app": "kube-proxy",
          "pod-template-generation": "1"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "kube-proxy",
            "uid": "b3c72459-1656-49c5-a068-32494f57f5b2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "kube-proxy",
            "configMap": {
              "name": "kube-proxy",
              "defaultMode": 420
            }
          },
          {
            "name": "xtables-lock",
            "hostPath": {
              "path": "/run/xtables.lock",
              "type": "FileOrCreate"
            }
          },
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "kube-proxy-token-wfz8m",
            "secret": {
              "secretName": "kube-proxy-token-wfz8m",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kube-proxy",
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "command": [
              "/usr/local/bin/kube-proxy",
              "--config=/var/lib/kube-proxy/config.conf",
              "--hostname-override=$(NODE_NAME)"
            ],
            "env": [
              {
                "name": "NODE_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-proxy",
                "mountPath": "/var/lib/kube-proxy"
              },
              {
                "name": "xtables-lock",
                "mountPath": "/run/xtables.lock"
              },
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "kube-proxy-token-wfz8m",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "kube-proxy",
        "serviceAccount": "kube-proxy",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-199.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priorityClassName": "system-node-critical",
        "priority": 2000001000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:31Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:33Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:33Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:31Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:32:31Z",
        "containerStatuses": [
          {
            "name": "kube-proxy",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:32:32Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-proxy@sha256:4ef8ca8fa1fbe311f3e6c5d6123d19f48a7bc62984ea8acfc8da8fbdac52bff7",
            "containerID": "docker://aaf325af66b010450a17b9fb1e53aa6d9692d453c323ff5f4e560990aaac8901"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kube-proxy-md9vq",
        "generateName": "kube-proxy-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-proxy-md9vq",
        "uid": "d490ec62-21eb-4089-9bef-5b2ecf73c2c6",
        "resourceVersion": "3250",
        "creationTimestamp": "2019-07-09T10:06:53Z",
        "labels": {
          "controller-revision-hash": "7bdbc788b8",
          "k8s-app": "kube-proxy",
          "pod-template-generation": "1"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "kube-proxy",
            "uid": "b3c72459-1656-49c5-a068-32494f57f5b2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "kube-proxy",
            "configMap": {
              "name": "kube-proxy",
              "defaultMode": 420
            }
          },
          {
            "name": "xtables-lock",
            "hostPath": {
              "path": "/run/xtables.lock",
              "type": "FileOrCreate"
            }
          },
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "kube-proxy-token-wfz8m",
            "secret": {
              "secretName": "kube-proxy-token-wfz8m",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kube-proxy",
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "command": [
              "/usr/local/bin/kube-proxy",
              "--config=/var/lib/kube-proxy/config.conf",
              "--hostname-override=$(NODE_NAME)"
            ],
            "env": [
              {
                "name": "NODE_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-proxy",
                "mountPath": "/var/lib/kube-proxy"
              },
              {
                "name": "xtables-lock",
                "mountPath": "/run/xtables.lock"
              },
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "kube-proxy-token-wfz8m",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "kube-proxy",
        "serviceAccount": "kube-proxy",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-203.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priorityClassName": "system-node-critical",
        "priority": 2000001000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:53Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:20Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:07:20Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:53Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "70.0.87.203",
        "startTime": "2019-07-09T10:06:53Z",
        "containerStatuses": [
          {
            "name": "kube-proxy",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:07:20Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-proxy@sha256:4ef8ca8fa1fbe311f3e6c5d6123d19f48a7bc62984ea8acfc8da8fbdac52bff7",
            "containerID": "docker://bf15c3c1a24d9fee79544ee4602e6dc874f2293ee644841e444c2745225f7765"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kube-proxy-p5pp2",
        "generateName": "kube-proxy-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-proxy-p5pp2",
        "uid": "41d55251-df81-40fe-8fa5-1456114ee978",
        "resourceVersion": "3033",
        "creationTimestamp": "2019-07-09T10:05:38Z",
        "labels": {
          "controller-revision-hash": "7bdbc788b8",
          "k8s-app": "kube-proxy",
          "pod-template-generation": "1"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "kube-proxy",
            "uid": "b3c72459-1656-49c5-a068-32494f57f5b2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "kube-proxy",
            "configMap": {
              "name": "kube-proxy",
              "defaultMode": 420
            }
          },
          {
            "name": "xtables-lock",
            "hostPath": {
              "path": "/run/xtables.lock",
              "type": "FileOrCreate"
            }
          },
          {
            "name": "lib-modules",
            "hostPath": {
              "path": "/lib/modules",
              "type": ""
            }
          },
          {
            "name": "kube-proxy-token-wfz8m",
            "secret": {
              "secretName": "kube-proxy-token-wfz8m",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "kube-proxy",
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "command": [
              "/usr/local/bin/kube-proxy",
              "--config=/var/lib/kube-proxy/config.conf",
              "--hostname-override=$(NODE_NAME)"
            ],
            "env": [
              {
                "name": "NODE_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "spec.nodeName"
                  }
                }
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-proxy",
                "mountPath": "/var/lib/kube-proxy"
              },
              {
                "name": "xtables-lock",
                "mountPath": "/run/xtables.lock"
              },
              {
                "name": "lib-modules",
                "readOnly": true,
                "mountPath": "/lib/modules"
              },
              {
                "name": "kube-proxy-token-wfz8m",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeSelector": {
          "beta.kubernetes.io/os": "linux"
        },
        "serviceAccountName": "kube-proxy",
        "serviceAccount": "kube-proxy",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-200.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "CriticalAddonsOnly",
            "operator": "Exists"
          },
          {
            "operator": "Exists"
          },
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priorityClassName": "system-node-critical",
        "priority": 2000001000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:05:38Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:07Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:06:07Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T10:05:38Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "70.0.87.200",
        "startTime": "2019-07-09T10:05:38Z",
        "containerStatuses": [
          {
            "name": "kube-proxy",
            "state": {
              "running": {
                "startedAt": "2019-07-09T10:06:07Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-proxy:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-proxy@sha256:4ef8ca8fa1fbe311f3e6c5d6123d19f48a7bc62984ea8acfc8da8fbdac52bff7",
            "containerID": "docker://d5fb0a1ee264a082ca4a6631abdb56aaa3beb85fecc7d3c71292f0180a04d54b"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "kube-scheduler-ip-70-0-87-199.brbnca.spcsdns.net",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/kube-scheduler-ip-70-0-87-199.brbnca.spcsdns.net",
        "uid": "d3234f3f-94fc-4b69-a7ae-7311f31d2c30",
        "resourceVersion": "461",
        "creationTimestamp": "2019-07-09T09:33:34Z",
        "labels": {
          "component": "kube-scheduler",
          "tier": "control-plane"
        },
        "annotations": {
          "kubernetes.io/config.hash": "31d9ee8b7fb12e797dc981a8686f6b2b",
          "kubernetes.io/config.mirror": "31d9ee8b7fb12e797dc981a8686f6b2b",
          "kubernetes.io/config.seen": "2019-07-09T09:31:54.302786259Z",
          "kubernetes.io/config.source": "file"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "kubeconfig",
            "hostPath": {
              "path": "/etc/kubernetes/scheduler.conf",
              "type": "FileOrCreate"
            }
          }
        ],
        "containers": [
          {
            "name": "kube-scheduler",
            "image": "k8s.gcr.io/kube-scheduler:v1.15.0",
            "command": [
              "kube-scheduler",
              "--bind-address=127.0.0.1",
              "--kubeconfig=/etc/kubernetes/scheduler.conf",
              "--leader-elect=true"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "kubeconfig",
                "readOnly": true,
                "mountPath": "/etc/kubernetes/scheduler.conf"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 8
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "nodeName": "ip-70-0-87-199.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "operator": "Exists",
            "effect": "NoExecute"
          }
        ],
        "priorityClassName": "system-cluster-critical",
        "priority": 2000000000,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:17Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-09T09:32:15Z"
          }
        ],
        "hostIP": "70.0.87.199",
        "podIP": "70.0.87.199",
        "startTime": "2019-07-09T09:32:15Z",
        "containerStatuses": [
          {
            "name": "kube-scheduler",
            "state": {
              "running": {
                "startedAt": "2019-07-09T09:32:15Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "k8s.gcr.io/kube-scheduler:v1.15.0",
            "imageID": "docker-pullable://k8s.gcr.io/kube-scheduler@sha256:591f79411b67ff0daebfb2bb7ff15c68f3d6b318d73f228f89a2fcb2c657ebea",
            "containerID": "docker://a4a571197b67a13158cff760c71fd6b9e728b8525e62253a7194a525ead1ca40"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "portworx-7dd86",
        "generateName": "portworx-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/portworx-7dd86",
        "uid": "e8ba2948-df45-4dd2-bb19-e431807be865",
        "resourceVersion": "2401738",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "controller-revision-hash": "68765dcdf",
          "name": "portworx",
          "pod-template-generation": "1"
        },
        "annotations": {
          "cluster-autoscaler.kubernetes.io/safe-to-evict": "false"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "portworx",
            "uid": "a16f2ab0-b0fd-44aa-982f-caed4ac962ea",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "diagsdump",
            "hostPath": {
              "path": "/var/cores",
              "type": ""
            }
          },
          {
            "name": "dockersock",
            "hostPath": {
              "path": "/var/run/docker.sock",
              "type": ""
            }
          },
          {
            "name": "containerdsock",
            "hostPath": {
              "path": "/run/containerd",
              "type": ""
            }
          },
          {
            "name": "etcpwx",
            "hostPath": {
              "path": "/etc/pwx",
              "type": ""
            }
          },
          {
            "name": "optpwx",
            "hostPath": {
              "path": "/opt/pwx",
              "type": ""
            }
          },
          {
            "name": "procmount",
            "hostPath": {
              "path": "/proc",
              "type": ""
            }
          },
          {
            "name": "sysdmount",
            "hostPath": {
              "path": "/etc/systemd/system",
              "type": ""
            }
          },
          {
            "name": "journalmount1",
            "hostPath": {
              "path": "/var/run/log",
              "type": ""
            }
          },
          {
            "name": "journalmount2",
            "hostPath": {
              "path": "/var/log",
              "type": ""
            }
          },
          {
            "name": "dbusmount",
            "hostPath": {
              "path": "/var/run/dbus",
              "type": ""
            }
          },
          {
            "name": "px-account-token-twplj",
            "secret": {
              "secretName": "px-account-token-twplj",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "portworx",
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "args": [
              "-k",
              "etcd:http://70.0.86.100:2379,etcd:http://70.0.86.92:2379,etcd:http://70.0.86.90:2379",
              "-c",
              "kkg-test-setup",
              "-a",
              "-secret_type",
              "k8s",
              "-x",
              "kubernetes",
              "-j",
              "auto"
            ],
            "env": [
              {
                "name": "AUTO_NODE_RECOVERY_TIMEOUT_IN_SECS",
                "value": "1500"
              },
              {
                "name": "PX_TEMPLATE_VERSION",
                "value": "v4"
              },
              {
                "name": "PX_IMAGE",
                "value": "portworx/px-base-enterprise:2.1.2"
              },
              {
                "name": "REGISTRY_USER",
                "value": "pwxbuild"
              },
              {
                "name": "REGISTRY_PASS",
                "value": "fridaydemos"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "diagsdump",
                "mountPath": "/var/cores"
              },
              {
                "name": "dockersock",
                "mountPath": "/var/run/docker.sock"
              },
              {
                "name": "containerdsock",
                "mountPath": "/run/containerd"
              },
              {
                "name": "etcpwx",
                "mountPath": "/etc/pwx"
              },
              {
                "name": "optpwx",
                "mountPath": "/opt/pwx"
              },
              {
                "name": "procmount",
                "mountPath": "/host_proc"
              },
              {
                "name": "sysdmount",
                "mountPath": "/etc/systemd/system"
              },
              {
                "name": "journalmount1",
                "readOnly": true,
                "mountPath": "/var/run/log"
              },
              {
                "name": "journalmount2",
                "readOnly": true,
                "mountPath": "/var/log"
              },
              {
                "name": "dbusmount",
                "mountPath": "/var/run/dbus"
              },
              {
                "name": "px-account-token-twplj",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/status",
                "port": 9001,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 840,
              "timeoutSeconds": 1,
              "periodSeconds": 30,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 9015,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/tmp/px-termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "px-account",
        "serviceAccount": "px-account",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-200.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:03Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:03Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "70.0.87.200",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "portworx",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:29:12Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "imageID": "docker-pullable://portworx/oci-monitor@sha256:630fb84a11f7dc7c4d6ab75ba96a6d098abc6a3dc1cff530edb85f45f49d06d7",
            "containerID": "docker://9e7cc57962481af8135e3e2aabf09bd378e47a59fddf811b438c9eef53cdf47a"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "portworx-bg69h",
        "generateName": "portworx-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/portworx-bg69h",
        "uid": "bd91ab80-5e7e-496d-8f11-b0303450da7e",
        "resourceVersion": "2401457",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "controller-revision-hash": "68765dcdf",
          "name": "portworx",
          "pod-template-generation": "1"
        },
        "annotations": {
          "cluster-autoscaler.kubernetes.io/safe-to-evict": "false"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "portworx",
            "uid": "a16f2ab0-b0fd-44aa-982f-caed4ac962ea",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "diagsdump",
            "hostPath": {
              "path": "/var/cores",
              "type": ""
            }
          },
          {
            "name": "dockersock",
            "hostPath": {
              "path": "/var/run/docker.sock",
              "type": ""
            }
          },
          {
            "name": "containerdsock",
            "hostPath": {
              "path": "/run/containerd",
              "type": ""
            }
          },
          {
            "name": "etcpwx",
            "hostPath": {
              "path": "/etc/pwx",
              "type": ""
            }
          },
          {
            "name": "optpwx",
            "hostPath": {
              "path": "/opt/pwx",
              "type": ""
            }
          },
          {
            "name": "procmount",
            "hostPath": {
              "path": "/proc",
              "type": ""
            }
          },
          {
            "name": "sysdmount",
            "hostPath": {
              "path": "/etc/systemd/system",
              "type": ""
            }
          },
          {
            "name": "journalmount1",
            "hostPath": {
              "path": "/var/run/log",
              "type": ""
            }
          },
          {
            "name": "journalmount2",
            "hostPath": {
              "path": "/var/log",
              "type": ""
            }
          },
          {
            "name": "dbusmount",
            "hostPath": {
              "path": "/var/run/dbus",
              "type": ""
            }
          },
          {
            "name": "px-account-token-twplj",
            "secret": {
              "secretName": "px-account-token-twplj",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "portworx",
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "args": [
              "-k",
              "etcd:http://70.0.86.100:2379,etcd:http://70.0.86.92:2379,etcd:http://70.0.86.90:2379",
              "-c",
              "kkg-test-setup",
              "-a",
              "-secret_type",
              "k8s",
              "-x",
              "kubernetes",
              "-j",
              "auto"
            ],
            "env": [
              {
                "name": "AUTO_NODE_RECOVERY_TIMEOUT_IN_SECS",
                "value": "1500"
              },
              {
                "name": "PX_TEMPLATE_VERSION",
                "value": "v4"
              },
              {
                "name": "PX_IMAGE",
                "value": "portworx/px-base-enterprise:2.1.2"
              },
              {
                "name": "REGISTRY_USER",
                "value": "pwxbuild"
              },
              {
                "name": "REGISTRY_PASS",
                "value": "fridaydemos"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "diagsdump",
                "mountPath": "/var/cores"
              },
              {
                "name": "dockersock",
                "mountPath": "/var/run/docker.sock"
              },
              {
                "name": "containerdsock",
                "mountPath": "/run/containerd"
              },
              {
                "name": "etcpwx",
                "mountPath": "/etc/pwx"
              },
              {
                "name": "optpwx",
                "mountPath": "/opt/pwx"
              },
              {
                "name": "procmount",
                "mountPath": "/host_proc"
              },
              {
                "name": "sysdmount",
                "mountPath": "/etc/systemd/system"
              },
              {
                "name": "journalmount1",
                "readOnly": true,
                "mountPath": "/var/run/log"
              },
              {
                "name": "journalmount2",
                "readOnly": true,
                "mountPath": "/var/log"
              },
              {
                "name": "dbusmount",
                "mountPath": "/var/run/dbus"
              },
              {
                "name": "px-account-token-twplj",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/status",
                "port": 9001,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 840,
              "timeoutSeconds": 1,
              "periodSeconds": 30,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 9015,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/tmp/px-termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "px-account",
        "serviceAccount": "px-account",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-233.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:36:30Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:36:30Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "70.0.87.233",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "portworx",
            "state": {
              "running": {
                "startedAt": "2019-07-25T12:37:47Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 2,
                "reason": "Error",
                "startedAt": "2019-07-25T08:29:17Z",
                "finishedAt": "2019-07-25T12:37:46Z",
                "containerID": "docker://ea8e4b018c58f80d28d6e66828dd41cde4cb840d50d29a5f02489ab3ede4f649"
              }
            },
            "ready": true,
            "restartCount": 1,
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "imageID": "docker-pullable://portworx/oci-monitor@sha256:630fb84a11f7dc7c4d6ab75ba96a6d098abc6a3dc1cff530edb85f45f49d06d7",
            "containerID": "docker://5d11afd40c06bf1fad85600896e190bebd83c41a8e3547228a28e0e6461a9d87"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "portworx-n6hfg",
        "generateName": "portworx-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/portworx-n6hfg",
        "uid": "9d8e790b-b40c-4fbb-863b-2660eb4bf021",
        "resourceVersion": "2401691",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "controller-revision-hash": "68765dcdf",
          "name": "portworx",
          "pod-template-generation": "1"
        },
        "annotations": {
          "cluster-autoscaler.kubernetes.io/safe-to-evict": "false"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "DaemonSet",
            "name": "portworx",
            "uid": "a16f2ab0-b0fd-44aa-982f-caed4ac962ea",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "diagsdump",
            "hostPath": {
              "path": "/var/cores",
              "type": ""
            }
          },
          {
            "name": "dockersock",
            "hostPath": {
              "path": "/var/run/docker.sock",
              "type": ""
            }
          },
          {
            "name": "containerdsock",
            "hostPath": {
              "path": "/run/containerd",
              "type": ""
            }
          },
          {
            "name": "etcpwx",
            "hostPath": {
              "path": "/etc/pwx",
              "type": ""
            }
          },
          {
            "name": "optpwx",
            "hostPath": {
              "path": "/opt/pwx",
              "type": ""
            }
          },
          {
            "name": "procmount",
            "hostPath": {
              "path": "/proc",
              "type": ""
            }
          },
          {
            "name": "sysdmount",
            "hostPath": {
              "path": "/etc/systemd/system",
              "type": ""
            }
          },
          {
            "name": "journalmount1",
            "hostPath": {
              "path": "/var/run/log",
              "type": ""
            }
          },
          {
            "name": "journalmount2",
            "hostPath": {
              "path": "/var/log",
              "type": ""
            }
          },
          {
            "name": "dbusmount",
            "hostPath": {
              "path": "/var/run/dbus",
              "type": ""
            }
          },
          {
            "name": "px-account-token-twplj",
            "secret": {
              "secretName": "px-account-token-twplj",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "portworx",
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "args": [
              "-k",
              "etcd:http://70.0.86.100:2379,etcd:http://70.0.86.92:2379,etcd:http://70.0.86.90:2379",
              "-c",
              "kkg-test-setup",
              "-a",
              "-secret_type",
              "k8s",
              "-x",
              "kubernetes",
              "-j",
              "auto"
            ],
            "env": [
              {
                "name": "AUTO_NODE_RECOVERY_TIMEOUT_IN_SECS",
                "value": "1500"
              },
              {
                "name": "PX_TEMPLATE_VERSION",
                "value": "v4"
              },
              {
                "name": "PX_IMAGE",
                "value": "portworx/px-base-enterprise:2.1.2"
              },
              {
                "name": "REGISTRY_USER",
                "value": "pwxbuild"
              },
              {
                "name": "REGISTRY_PASS",
                "value": "fridaydemos"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "diagsdump",
                "mountPath": "/var/cores"
              },
              {
                "name": "dockersock",
                "mountPath": "/var/run/docker.sock"
              },
              {
                "name": "containerdsock",
                "mountPath": "/run/containerd"
              },
              {
                "name": "etcpwx",
                "mountPath": "/etc/pwx"
              },
              {
                "name": "optpwx",
                "mountPath": "/opt/pwx"
              },
              {
                "name": "procmount",
                "mountPath": "/host_proc"
              },
              {
                "name": "sysdmount",
                "mountPath": "/etc/systemd/system"
              },
              {
                "name": "journalmount1",
                "readOnly": true,
                "mountPath": "/var/run/log"
              },
              {
                "name": "journalmount2",
                "readOnly": true,
                "mountPath": "/var/log"
              },
              {
                "name": "dbusmount",
                "mountPath": "/var/run/dbus"
              },
              {
                "name": "px-account-token-twplj",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/status",
                "port": 9001,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 840,
              "timeoutSeconds": 1,
              "periodSeconds": 30,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/health",
                "port": 9015,
                "host": "127.0.0.1",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/tmp/px-termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always",
            "securityContext": {
              "privileged": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "px-account",
        "serviceAccount": "px-account",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "hostNetwork": true,
        "securityContext": {},
        "affinity": {
          "nodeAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": {
              "nodeSelectorTerms": [
                {
                  "matchFields": [
                    {
                      "key": "metadata.name",
                      "operator": "In",
                      "values": [
                        "ip-70-0-87-203.brbnca.spcsdns.net"
                      ]
                    }
                  ]
                }
              ]
            }
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute"
          },
          {
            "key": "node.kubernetes.io/disk-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/memory-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/pid-pressure",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/unschedulable",
            "operator": "Exists",
            "effect": "NoSchedule"
          },
          {
            "key": "node.kubernetes.io/network-unavailable",
            "operator": "Exists",
            "effect": "NoSchedule"
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:37:56Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:37:56Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "70.0.87.203",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "portworx",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:29:12Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "portworx/oci-monitor:2.1.2-rc5",
            "imageID": "docker-pullable://portworx/oci-monitor@sha256:630fb84a11f7dc7c4d6ab75ba96a6d098abc6a3dc1cff530edb85f45f49d06d7",
            "containerID": "docker://1f63e0bce399968ac7529e52245abbdd4eaf71b6732882891970c481f0aff06d"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "prometheus-operator-55bd6468b9-54t8b",
        "generateName": "prometheus-operator-55bd6468b9-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/prometheus-operator-55bd6468b9-54t8b",
        "uid": "cabec025-53fe-45a2-8eb8-cb50ecca423e",
        "resourceVersion": "2160106",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "k8s-app": "prometheus-operator",
          "pod-template-hash": "55bd6468b9"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "prometheus-operator-55bd6468b9",
            "uid": "1c9b1ab6-19ba-445d-af24-8831da96c848",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "prometheus-operator-token-svgzh",
            "secret": {
              "secretName": "prometheus-operator-token-svgzh",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "prometheus-operator",
            "image": "quay.io/coreos/prometheus-operator:v0.29.0",
            "args": [
              "--kubelet-service=kube-system/kubelet",
              "--config-reloader-image=quay.io/coreos/configmap-reload:v0.0.1"
            ],
            "ports": [
              {
                "name": "http",
                "containerPort": 8080,
                "protocol": "TCP"
              }
            ],
            "resources": {
              "limits": {
                "cpu": "200m",
                "memory": "100Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "50Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "prometheus-operator-token-svgzh",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "prometheus-operator",
        "serviceAccount": "prometheus-operator",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {
          "runAsUser": 65534,
          "runAsNonRoot": true
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:31:47Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:31:47Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.133",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "prometheus-operator",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:31:46Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/prometheus-operator:v0.29.0",
            "imageID": "docker-pullable://quay.io/coreos/prometheus-operator@sha256:5abe9bdfd93ac22954e3281315637d9721d66539134e1c7ed4e97f13819e62f7",
            "containerID": "docker://27a9a271be938eaa6b3e4449e6d6265282cd44fc67939f5ddd9737afae729ed8"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "prometheus-prometheus-0",
        "generateName": "prometheus-prometheus-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/prometheus-prometheus-0",
        "uid": "caeaf055-e661-4494-b415-d250d92f6013",
        "resourceVersion": "2161557",
        "creationTimestamp": "2019-07-25T08:38:49Z",
        "labels": {
          "app": "prometheus",
          "controller-revision-hash": "prometheus-prometheus-59bc8994d5",
          "prometheus": "prometheus",
          "statefulset.kubernetes.io/pod-name": "prometheus-prometheus-0"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "StatefulSet",
            "name": "prometheus-prometheus",
            "uid": "baea3b0f-1e40-4915-9db2-fd4fbdcf7d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config",
            "secret": {
              "secretName": "prometheus-prometheus",
              "defaultMode": 420
            }
          },
          {
            "name": "config-out",
            "emptyDir": {}
          },
          {
            "name": "prometheus-prometheus-rulefiles-0",
            "configMap": {
              "name": "prometheus-prometheus-rulefiles-0",
              "defaultMode": 420
            }
          },
          {
            "name": "prometheus-prometheus-db",
            "emptyDir": {}
          },
          {
            "name": "prometheus-token-6bbmp",
            "secret": {
              "secretName": "prometheus-token-6bbmp",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "prometheus",
            "image": "quay.io/prometheus/prometheus:v2.7.1",
            "args": [
              "--web.console.templates=/etc/prometheus/consoles",
              "--web.console.libraries=/etc/prometheus/console_libraries",
              "--config.file=/etc/prometheus/config_out/prometheus.env.yaml",
              "--storage.tsdb.path=/prometheus",
              "--storage.tsdb.retention=24h",
              "--web.enable-lifecycle",
              "--storage.tsdb.no-lockfile",
              "--web.route-prefix=/",
              "--log.level=debug"
            ],
            "ports": [
              {
                "name": "web",
                "containerPort": 9090,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "config-out",
                "readOnly": true,
                "mountPath": "/etc/prometheus/config_out"
              },
              {
                "name": "prometheus-prometheus-db",
                "mountPath": "/prometheus"
              },
              {
                "name": "prometheus-prometheus-rulefiles-0",
                "mountPath": "/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/-/healthy",
                "port": "web",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 3,
              "periodSeconds": 5,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/-/ready",
                "port": "web",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 3,
              "periodSeconds": 5,
              "successThreshold": 1,
              "failureThreshold": 120
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "prometheus-config-reloader",
            "image": "quay.io/coreos/prometheus-config-reloader:v0.29.0",
            "command": [
              "/bin/prometheus-config-reloader"
            ],
            "args": [
              "--log-format=logfmt",
              "--reload-url=http://localhost:9090/-/reload",
              "--config-file=/etc/prometheus/config/prometheus.yaml.gz",
              "--config-envsubst-file=/etc/prometheus/config_out/prometheus.env.yaml"
            ],
            "env": [
              {
                "name": "POD_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "metadata.name"
                  }
                }
              }
            ],
            "resources": {
              "limits": {
                "cpu": "50m",
                "memory": "50Mi"
              },
              "requests": {
                "cpu": "50m",
                "memory": "50Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config",
                "mountPath": "/etc/prometheus/config"
              },
              {
                "name": "config-out",
                "mountPath": "/etc/prometheus/config_out"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "rules-configmap-reloader",
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "args": [
              "--webhook-url=http://localhost:9090/-/reload",
              "--volume-dir=/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
            ],
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "25Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "25Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "prometheus-prometheus-rulefiles-0",
                "mountPath": "/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 600,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "prometheus",
        "serviceAccount": "prometheus",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "hostname": "prometheus-prometheus-0",
        "subdomain": "prometheus-operated",
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:48Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:48Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:49Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.200",
        "startTime": "2019-07-25T08:38:59Z",
        "containerStatuses": [
          {
            "name": "prometheus",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:40:45Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 1,
                "reason": "Error",
                "startedAt": "2019-07-25T08:40:31Z",
                "finishedAt": "2019-07-25T08:40:31Z",
                "containerID": "docker://409543df534b101f2269f508ae775cd5bcc42b835c0895dd4df966443942f9de"
              }
            },
            "ready": true,
            "restartCount": 1,
            "image": "quay.io/prometheus/prometheus:v2.7.1",
            "imageID": "docker-pullable://quay.io/prometheus/prometheus@sha256:f21aa9225a1ba630ea1671c4c7d8595aa87a6e89bbb6525c56703ed317ae9226",
            "containerID": "docker://0de31503c342e2d77f7583723e1cb6217624d2ce70b6d21054fddc849f680678"
          },
          {
            "name": "prometheus-config-reloader",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:40:34Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/prometheus-config-reloader:v0.29.0",
            "imageID": "docker-pullable://quay.io/coreos/prometheus-config-reloader@sha256:d4b5c90b0937b568d25baa38d8ca61ac927069f30f02e8757f0eae9ee335be23",
            "containerID": "docker://48c8403f19a44275e8896517d8fbd7b7c716e838436e63495c68afa6340ed292"
          },
          {
            "name": "rules-configmap-reloader",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:40:44Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "imageID": "docker-pullable://quay.io/coreos/configmap-reload@sha256:e2fd60ff0ae4500a75b80ebaa30e0e7deba9ad107833e8ca53f0047c42c5a057",
            "containerID": "docker://dd57c9aa529af0b07a284f96d906e98e37457b64e13e8c0d09a3c6e3e38b553d"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "prometheus-prometheus-1",
        "generateName": "prometheus-prometheus-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/prometheus-prometheus-1",
        "uid": "8f8513b4-e049-47d9-9ba0-c2096a0db433",
        "resourceVersion": "2161290",
        "creationTimestamp": "2019-07-25T08:38:49Z",
        "labels": {
          "app": "prometheus",
          "controller-revision-hash": "prometheus-prometheus-59bc8994d5",
          "prometheus": "prometheus",
          "statefulset.kubernetes.io/pod-name": "prometheus-prometheus-1"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "StatefulSet",
            "name": "prometheus-prometheus",
            "uid": "baea3b0f-1e40-4915-9db2-fd4fbdcf7d02",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config",
            "secret": {
              "secretName": "prometheus-prometheus",
              "defaultMode": 420
            }
          },
          {
            "name": "config-out",
            "emptyDir": {}
          },
          {
            "name": "prometheus-prometheus-rulefiles-0",
            "configMap": {
              "name": "prometheus-prometheus-rulefiles-0",
              "defaultMode": 420
            }
          },
          {
            "name": "prometheus-prometheus-db",
            "emptyDir": {}
          },
          {
            "name": "prometheus-token-6bbmp",
            "secret": {
              "secretName": "prometheus-token-6bbmp",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "prometheus",
            "image": "quay.io/prometheus/prometheus:v2.7.1",
            "args": [
              "--web.console.templates=/etc/prometheus/consoles",
              "--web.console.libraries=/etc/prometheus/console_libraries",
              "--config.file=/etc/prometheus/config_out/prometheus.env.yaml",
              "--storage.tsdb.path=/prometheus",
              "--storage.tsdb.retention=24h",
              "--web.enable-lifecycle",
              "--storage.tsdb.no-lockfile",
              "--web.route-prefix=/",
              "--log.level=debug"
            ],
            "ports": [
              {
                "name": "web",
                "containerPort": 9090,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "config-out",
                "readOnly": true,
                "mountPath": "/etc/prometheus/config_out"
              },
              {
                "name": "prometheus-prometheus-db",
                "mountPath": "/prometheus"
              },
              {
                "name": "prometheus-prometheus-rulefiles-0",
                "mountPath": "/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/-/healthy",
                "port": "web",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 3,
              "periodSeconds": 5,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/-/ready",
                "port": "web",
                "scheme": "HTTP"
              },
              "timeoutSeconds": 3,
              "periodSeconds": 5,
              "successThreshold": 1,
              "failureThreshold": 120
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "prometheus-config-reloader",
            "image": "quay.io/coreos/prometheus-config-reloader:v0.29.0",
            "command": [
              "/bin/prometheus-config-reloader"
            ],
            "args": [
              "--log-format=logfmt",
              "--reload-url=http://localhost:9090/-/reload",
              "--config-file=/etc/prometheus/config/prometheus.yaml.gz",
              "--config-envsubst-file=/etc/prometheus/config_out/prometheus.env.yaml"
            ],
            "env": [
              {
                "name": "POD_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "metadata.name"
                  }
                }
              }
            ],
            "resources": {
              "limits": {
                "cpu": "50m",
                "memory": "50Mi"
              },
              "requests": {
                "cpu": "50m",
                "memory": "50Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "config",
                "mountPath": "/etc/prometheus/config"
              },
              {
                "name": "config-out",
                "mountPath": "/etc/prometheus/config_out"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          },
          {
            "name": "rules-configmap-reloader",
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "args": [
              "--webhook-url=http://localhost:9090/-/reload",
              "--volume-dir=/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
            ],
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "25Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "25Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "prometheus-prometheus-rulefiles-0",
                "mountPath": "/etc/prometheus/rules/prometheus-prometheus-rulefiles-0"
              },
              {
                "name": "prometheus-token-6bbmp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 600,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "prometheus",
        "serviceAccount": "prometheus",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "hostname": "prometheus-prometheus-1",
        "subdomain": "prometheus-operated",
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:49Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:39:24Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:39:24Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:49Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.134",
        "startTime": "2019-07-25T08:38:49Z",
        "containerStatuses": [
          {
            "name": "prometheus",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:39:23Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 1,
                "reason": "Error",
                "startedAt": "2019-07-25T08:39:02Z",
                "finishedAt": "2019-07-25T08:39:02Z",
                "containerID": "docker://fae4b3748bd54a0ec2e10e75e16be91adb46e4f6800be7bee1854b93c1065549"
              }
            },
            "ready": true,
            "restartCount": 1,
            "image": "quay.io/prometheus/prometheus:v2.7.1",
            "imageID": "docker-pullable://quay.io/prometheus/prometheus@sha256:f21aa9225a1ba630ea1671c4c7d8595aa87a6e89bbb6525c56703ed317ae9226",
            "containerID": "docker://0c7ee341b03d752e75e48f76fce56d1ffb186784e8290a73c16e42101ee1ec34"
          },
          {
            "name": "prometheus-config-reloader",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:39:20Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/prometheus-config-reloader:v0.29.0",
            "imageID": "docker-pullable://quay.io/coreos/prometheus-config-reloader@sha256:d4b5c90b0937b568d25baa38d8ca61ac927069f30f02e8757f0eae9ee335be23",
            "containerID": "docker://445b92036d0f8c6abf3afb5d19779c20de5833d6f16f508fe2ca3cdd0db3efc5"
          },
          {
            "name": "rules-configmap-reloader",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:39:22Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "quay.io/coreos/configmap-reload:v0.0.1",
            "imageID": "docker-pullable://quay.io/coreos/configmap-reload@sha256:e2fd60ff0ae4500a75b80ebaa30e0e7deba9ad107833e8ca53f0047c42c5a057",
            "containerID": "docker://9fefdc400115ee087e732c0eb2fe008f649f366e24184a0446236a641a5d1ba2"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "px-lighthouse-56dccdd987-t4qzf",
        "generateName": "px-lighthouse-56dccdd987-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/px-lighthouse-56dccdd987-t4qzf",
        "uid": "4e0494a5-3e35-4d78-ad9a-9e4a86d188b9",
        "resourceVersion": "2162741",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "pod-template-hash": "56dccdd987",
          "tier": "px-web-console"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "px-lighthouse-56dccdd987",
            "uid": "bb23e2d5-40de-4902-bd23-1de65e85f4e7",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "config",
            "emptyDir": {}
          },
          {
            "name": "px-lh-account-token-mghqf",
            "secret": {
              "secretName": "px-lh-account-token-mghqf",
              "defaultMode": 420
            }
          }
        ],
        "initContainers": [
          {
            "name": "config-init",
            "image": "portworx/lh-config-sync:0.3",
            "args": [
              "init"
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "config",
                "mountPath": "/config/lh"
              },
              {
                "name": "px-lh-account-token-mghqf",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "containers": [
          {
            "name": "px-lighthouse",
            "image": "portworx/px-lighthouse:2.0.4",
            "args": [
              "-kubernetes",
              "true"
            ],
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              },
              {
                "containerPort": 443,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "config",
                "mountPath": "/config/lh"
              },
              {
                "name": "px-lh-account-token-mghqf",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          },
          {
            "name": "config-sync",
            "image": "portworx/lh-config-sync:0.3",
            "args": [
              "sync"
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "config",
                "mountPath": "/config/lh"
              },
              {
                "name": "px-lh-account-token-mghqf",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          },
          {
            "name": "stork-connector",
            "image": "portworx/lh-stork-connector:0.1",
            "resources": {},
            "volumeMounts": [
              {
                "name": "px-lh-account-token-mghqf",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "px-lh-account",
        "serviceAccount": "px-lh-account",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:47:21Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:48:17Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:48:17Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "192.168.220.196",
        "startTime": "2019-07-25T08:29:00Z",
        "initContainerStatuses": [
          {
            "name": "config-init",
            "state": {
              "terminated": {
                "exitCode": 0,
                "reason": "Completed",
                "startedAt": "2019-07-25T08:47:19Z",
                "finishedAt": "2019-07-25T08:47:20Z",
                "containerID": "docker://6da5d14671e132f3092e478aae046d3e276a50697307df343ab0e0c4aa040e70"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 5,
            "image": "portworx/lh-config-sync:0.3",
            "imageID": "docker-pullable://portworx/lh-config-sync@sha256:b98e5f38779f05917ae628bc86ffc411c69b92354526cd9d598ae2ec02d552c5",
            "containerID": "docker://6da5d14671e132f3092e478aae046d3e276a50697307df343ab0e0c4aa040e70"
          }
        ],
        "containerStatuses": [
          {
            "name": "config-sync",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:47:54Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "portworx/lh-config-sync:0.3",
            "imageID": "docker-pullable://portworx/lh-config-sync@sha256:b98e5f38779f05917ae628bc86ffc411c69b92354526cd9d598ae2ec02d552c5",
            "containerID": "docker://3beec24724f80235ff29a234daa8ac77c2a7f8795920be1b90161368ec0f0ee2"
          },
          {
            "name": "px-lighthouse",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:47:54Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "portworx/px-lighthouse:2.0.4",
            "imageID": "docker-pullable://portworx/px-lighthouse@sha256:7933c904f7f6a2e2f460d4870ee9cc27aa9daf050494c96099e5a7775afe578b",
            "containerID": "docker://9a23b1f956f0a4b4771608b1806dacca714b00bb50a7f9ae409f4771be2c002e"
          },
          {
            "name": "stork-connector",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:48:17Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "portworx/lh-stork-connector:0.1",
            "imageID": "docker-pullable://portworx/lh-stork-connector@sha256:68f99733225fec216144a3c21ba22113e59f7217037858599393fa2886f53fae",
            "containerID": "docker://1e87a29c292a3e2665e61a3043ef5deb375c69b7cf7d702708e9dd3db13440b0"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "stork-79476fb994-b2bkw",
        "generateName": "stork-79476fb994-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-79476fb994-b2bkw",
        "uid": "a9b23e60-345c-4a02-a490-981c50421976",
        "resourceVersion": "2161178",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "name": "stork",
          "pod-template-hash": "79476fb994",
          "tier": "control-plane"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-79476fb994",
            "uid": "8276bedb-387a-4a88-9d25-52413393eaf7",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-account-token-dl79l",
            "secret": {
              "secretName": "stork-account-token-dl79l",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork",
            "image": "openstorage/stork:2.1.2",
            "command": [
              "/stork",
              "--driver=pxd",
              "--verbose",
              "--leader-elect=true",
              "--health-monitor-interval=120"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-account-token-dl79l",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-account",
        "serviceAccount": "stork-account",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:38:59Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "192.168.220.197",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:38:59Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "openstorage/stork:2.1.2",
            "imageID": "docker-pullable://openstorage/stork@sha256:d15011d17fb7b854c3093131636ceed3957fd54f38d7379536ef64f1b5d39be9",
            "containerID": "docker://aa1ac0ddf961ef101509e4969ad83d2054fef9715c8680eaed900b35fab9fbd1"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "stork-79476fb994-rxr7w",
        "generateName": "stork-79476fb994-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-79476fb994-rxr7w",
        "uid": "fed6610e-12b9-4773-b8dc-24e56ad41ef5",
        "resourceVersion": "2161459",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "name": "stork",
          "pod-template-hash": "79476fb994",
          "tier": "control-plane"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-79476fb994",
            "uid": "8276bedb-387a-4a88-9d25-52413393eaf7",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-account-token-dl79l",
            "secret": {
              "secretName": "stork-account-token-dl79l",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork",
            "image": "openstorage/stork:2.1.2",
            "command": [
              "/stork",
              "--driver=pxd",
              "--verbose",
              "--leader-elect=true",
              "--health-monitor-interval=120"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-account-token-dl79l",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-account",
        "serviceAccount": "stork-account",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:23Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:40:23Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.198",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:40:22Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "openstorage/stork:2.1.2",
            "imageID": "docker-pullable://openstorage/stork@sha256:d15011d17fb7b854c3093131636ceed3957fd54f38d7379536ef64f1b5d39be9",
            "containerID": "docker://e946355e4bf10ae4dc5e60a59a691301053279a782e72707f9644b54b4f7cae6"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "stork-79476fb994-wgg66",
        "generateName": "stork-79476fb994-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-79476fb994-wgg66",
        "uid": "e2c81f49-be19-4bdf-860f-2df2655c0274",
        "resourceVersion": "2161388",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "name": "stork",
          "pod-template-hash": "79476fb994",
          "tier": "control-plane"
        },
        "annotations": {
          "scheduler.alpha.kubernetes.io/critical-pod": ""
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-79476fb994",
            "uid": "8276bedb-387a-4a88-9d25-52413393eaf7",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-account-token-dl79l",
            "secret": {
              "secretName": "stork-account-token-dl79l",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork",
            "image": "openstorage/stork:2.1.2",
            "command": [
              "/stork",
              "--driver=pxd",
              "--verbose",
              "--leader-elect=true",
              "--health-monitor-interval=120"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-account-token-dl79l",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-account",
        "serviceAccount": "stork-account",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:39:58Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:39:58Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.132",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:39:57Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 2,
                "reason": "Error",
                "startedAt": "2019-07-25T08:39:20Z",
                "finishedAt": "2019-07-25T08:39:41Z",
                "containerID": "docker://b7d1c006509efb21e93394e9ca4c2a35b43a10951e95e275d5850b7365883b22"
              }
            },
            "ready": true,
            "restartCount": 2,
            "image": "openstorage/stork:2.1.2",
            "imageID": "docker-pullable://openstorage/stork@sha256:d15011d17fb7b854c3093131636ceed3957fd54f38d7379536ef64f1b5d39be9",
            "containerID": "docker://880a5037283229f6f6a97047c348b2b49b95a4ecdf1a6a3142fd890489234a81"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "stork-scheduler-86b4b5c55d-hmlgn",
        "generateName": "stork-scheduler-86b4b5c55d-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-scheduler-86b4b5c55d-hmlgn",
        "uid": "d1c34194-221c-4097-890f-c6aa4947ed33",
        "resourceVersion": "2160016",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "component": "scheduler",
          "pod-template-hash": "86b4b5c55d",
          "tier": "control-plane"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-scheduler-86b4b5c55d",
            "uid": "3ae70f2c-d637-4319-ac9f-740d316a65d2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-scheduler-account-token-qf9mp",
            "secret": {
              "secretName": "stork-scheduler-account-token-qf9mp",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork-scheduler",
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "command": [
              "/usr/local/bin/kube-scheduler",
              "--address=0.0.0.0",
              "--leader-elect=true",
              "--scheduler-name=stork",
              "--policy-configmap=stork-config",
              "--policy-configmap-namespace=kube-system",
              "--lock-object-name=stork-scheduler"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-scheduler-account-token-qf9mp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-scheduler-account",
        "serviceAccount": "stork-scheduler-account",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork-scheduler"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:31:19Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:31:19Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.131",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork-scheduler",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:30:13Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "imageID": "docker-pullable://gcr.io/google_containers/kube-scheduler-amd64@sha256:c6872392573b95225ef37175530b2b81c81a14641f2d630eec3f5452e3244e3c",
            "containerID": "docker://d5881f9dd782d5dc45684c4054541bff8e8f391bc1ae58e4100bf701720d9718"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "stork-scheduler-86b4b5c55d-l89sj",
        "generateName": "stork-scheduler-86b4b5c55d-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-scheduler-86b4b5c55d-l89sj",
        "uid": "c1b0f3d3-8e51-4ef2-a50b-cc56aee41aac",
        "resourceVersion": "2159735",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "component": "scheduler",
          "pod-template-hash": "86b4b5c55d",
          "tier": "control-plane"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-scheduler-86b4b5c55d",
            "uid": "3ae70f2c-d637-4319-ac9f-740d316a65d2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-scheduler-account-token-qf9mp",
            "secret": {
              "secretName": "stork-scheduler-account-token-qf9mp",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork-scheduler",
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "command": [
              "/usr/local/bin/kube-scheduler",
              "--address=0.0.0.0",
              "--leader-elect=true",
              "--scheduler-name=stork",
              "--policy-configmap=stork-config",
              "--policy-configmap-namespace=kube-system",
              "--lock-object-name=stork-scheduler"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-scheduler-account-token-qf9mp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-scheduler-account",
        "serviceAccount": "stork-scheduler-account",
        "nodeName": "ip-70-0-87-203.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork-scheduler"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:24Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:24Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.203",
        "podIP": "192.168.220.194",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork-scheduler",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:29:19Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "imageID": "docker-pullable://gcr.io/google_containers/kube-scheduler-amd64@sha256:c6872392573b95225ef37175530b2b81c81a14641f2d630eec3f5452e3244e3c",
            "containerID": "docker://3b910a46f9c7247bb08cea05443db979b4ead6ba029a83d93d1c15cd6fcedd3e"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "stork-scheduler-86b4b5c55d-qj8kr",
        "generateName": "stork-scheduler-86b4b5c55d-",
        "namespace": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system/pods/stork-scheduler-86b4b5c55d-qj8kr",
        "uid": "91fe7689-9693-400a-a89f-dbcd473fd377",
        "resourceVersion": "2160680",
        "creationTimestamp": "2019-07-25T08:29:00Z",
        "labels": {
          "component": "scheduler",
          "pod-template-hash": "86b4b5c55d",
          "tier": "control-plane"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "stork-scheduler-86b4b5c55d",
            "uid": "3ae70f2c-d637-4319-ac9f-740d316a65d2",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "stork-scheduler-account-token-qf9mp",
            "secret": {
              "secretName": "stork-scheduler-account-token-qf9mp",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "stork-scheduler",
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "command": [
              "/usr/local/bin/kube-scheduler",
              "--address=0.0.0.0",
              "--leader-elect=true",
              "--scheduler-name=stork",
              "--policy-configmap=stork-config",
              "--policy-configmap-namespace=kube-system",
              "--lock-object-name=stork-scheduler"
            ],
            "resources": {
              "requests": {
                "cpu": "100m"
              }
            },
            "volumeMounts": [
              {
                "name": "stork-scheduler-account-token-qf9mp",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 15,
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/healthz",
                "port": 10251,
                "scheme": "HTTP"
              },
              "timeoutSeconds": 1,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "stork-scheduler-account",
        "serviceAccount": "stork-scheduler-account",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "affinity": {
          "podAntiAffinity": {
            "requiredDuringSchedulingIgnoredDuringExecution": [
              {
                "labelSelector": {
                  "matchExpressions": [
                    {
                      "key": "name",
                      "operator": "In",
                      "values": [
                        "stork-scheduler"
                      ]
                    }
                  ]
                },
                "topologyKey": "kubernetes.io/hostname"
              }
            ]
          }
        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:35:36Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:35:36Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:29:00Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.197",
        "startTime": "2019-07-25T08:29:00Z",
        "containerStatuses": [
          {
            "name": "stork-scheduler",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:35:27Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "gcr.io/google_containers/kube-scheduler-amd64:v1.13.1",
            "imageID": "docker-pullable://gcr.io/google_containers/kube-scheduler-amd64@sha256:c6872392573b95225ef37175530b2b81c81a14641f2d630eec3f5452e3244e3c",
            "containerID": "docker://6948fe2d6587a405c165be136ea9db9f0bfed6246b13610975034b0f056a8149"
          }
        ],
        "qosClass": "Burstable"
      }
    },
    {
      "metadata": {
        "name": "wordpress-7f6d665c6f-5wpm6",
        "generateName": "wordpress-7f6d665c6f-",
        "namespace": "wp1",
        "selfLink": "/api/v1/namespaces/wp1/pods/wordpress-7f6d665c6f-5wpm6",
        "uid": "a339e976-8271-4ab3-a533-8a240cbee8a9",
        "resourceVersion": "2401727",
        "creationTimestamp": "2019-07-25T08:49:59Z",
        "labels": {
          "app": "wordpress",
          "pod-template-hash": "7f6d665c6f",
          "tier": "frontend"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "wordpress-7f6d665c6f",
            "uid": "8b95963d-0032-45af-8c05-12605df36a43",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "wordpress-persistent-storage",
            "persistentVolumeClaim": {
              "claimName": "wp-pv-claim"
            }
          },
          {
            "name": "default-token-7gh6x",
            "secret": {
              "secretName": "default-token-7gh6x",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "wordpress",
            "image": "wordpress:4.8-apache",
            "ports": [
              {
                "name": "wordpress",
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "env": [
              {
                "name": "WORDPRESS_DB_HOST",
                "value": "wordpress-mysql"
              },
              {
                "name": "WORDPRESS_DB_PASSWORD",
                "value": "mysql@123"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "wordpress-persistent-storage",
                "mountPath": "/var/www/html"
              },
              {
                "name": "default-token-7gh6x",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "stork",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:02Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:02Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:02Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:02Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.136",
        "startTime": "2019-07-25T08:50:02Z",
        "containerStatuses": [
          {
            "name": "wordpress",
            "state": {
              "running": {
                "startedAt": "2019-07-26T10:38:01Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 137,
                "reason": "Error",
                "startedAt": "2019-07-25T08:53:12Z",
                "finishedAt": "2019-07-26T10:38:00Z",
                "containerID": "docker://b9848aab314b3b4a46ab9da781af5df528d52b3c8f2671444bccd030533cf5de"
              }
            },
            "ready": true,
            "restartCount": 2,
            "image": "wordpress:4.8-apache",
            "imageID": "docker-pullable://wordpress@sha256:6216f64ab88fc51d311e38c7f69ca3f9aaba621492b4f1fa93ddf63093768845",
            "containerID": "docker://6fe07f0324d9921684602200f6ea0b5b2dc51941e5fbfb9cab4cb4623ac8104b"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "wordpress-7f6d665c6f-7qcch",
        "generateName": "wordpress-7f6d665c6f-",
        "namespace": "wp1",
        "selfLink": "/api/v1/namespaces/wp1/pods/wordpress-7f6d665c6f-7qcch",
        "uid": "4531980a-32a3-46e4-9ded-ba81f0ae13f7",
        "resourceVersion": "2401724",
        "creationTimestamp": "2019-07-25T08:49:59Z",
        "labels": {
          "app": "wordpress",
          "pod-template-hash": "7f6d665c6f",
          "tier": "frontend"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "wordpress-7f6d665c6f",
            "uid": "8b95963d-0032-45af-8c05-12605df36a43",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "wordpress-persistent-storage",
            "persistentVolumeClaim": {
              "claimName": "wp-pv-claim"
            }
          },
          {
            "name": "default-token-7gh6x",
            "secret": {
              "secretName": "default-token-7gh6x",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "wordpress",
            "image": "wordpress:4.8-apache",
            "ports": [
              {
                "name": "wordpress",
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "env": [
              {
                "name": "WORDPRESS_DB_HOST",
                "value": "wordpress-mysql"
              },
              {
                "name": "WORDPRESS_DB_PASSWORD",
                "value": "mysql@123"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "wordpress-persistent-storage",
                "mountPath": "/var/www/html"
              },
              {
                "name": "default-token-7gh6x",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "stork",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:01Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:02Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:38:02Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:01Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.135",
        "startTime": "2019-07-25T08:50:01Z",
        "containerStatuses": [
          {
            "name": "wordpress",
            "state": {
              "running": {
                "startedAt": "2019-07-26T10:38:01Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 137,
                "reason": "Error",
                "startedAt": "2019-07-25T08:53:05Z",
                "finishedAt": "2019-07-26T10:38:00Z",
                "containerID": "docker://2865f2f738808f3dfd87ff933f80af03576cd12667e0c74023701f83934dd0ad"
              }
            },
            "ready": true,
            "restartCount": 2,
            "image": "wordpress:4.8-apache",
            "imageID": "docker-pullable://wordpress@sha256:6216f64ab88fc51d311e38c7f69ca3f9aaba621492b4f1fa93ddf63093768845",
            "containerID": "docker://16a13594bb284230ea89c40a6752e36866386510c964ccdd8191628e3b65de7e"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "wordpress-7f6d665c6f-ddjj6",
        "generateName": "wordpress-7f6d665c6f-",
        "namespace": "wp1",
        "selfLink": "/api/v1/namespaces/wp1/pods/wordpress-7f6d665c6f-ddjj6",
        "uid": "526ce73d-c5ca-44db-a6a1-457811e0319e",
        "resourceVersion": "2401440",
        "creationTimestamp": "2019-07-25T08:49:59Z",
        "labels": {
          "app": "wordpress",
          "pod-template-hash": "7f6d665c6f",
          "tier": "frontend"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "wordpress-7f6d665c6f",
            "uid": "8b95963d-0032-45af-8c05-12605df36a43",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "wordpress-persistent-storage",
            "persistentVolumeClaim": {
              "claimName": "wp-pv-claim"
            }
          },
          {
            "name": "default-token-7gh6x",
            "secret": {
              "secretName": "default-token-7gh6x",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "wordpress",
            "image": "wordpress:4.8-apache",
            "ports": [
              {
                "name": "wordpress",
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "env": [
              {
                "name": "WORDPRESS_DB_HOST",
                "value": "wordpress-mysql"
              },
              {
                "name": "WORDPRESS_DB_PASSWORD",
                "value": "mysql@123"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "wordpress-persistent-storage",
                "mountPath": "/var/www/html"
              },
              {
                "name": "default-token-7gh6x",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-233.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "stork",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:01Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:36:25Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-26T10:36:25Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:01Z"
          }
        ],
        "hostIP": "70.0.87.233",
        "podIP": "192.168.40.201",
        "startTime": "2019-07-25T08:50:01Z",
        "containerStatuses": [
          {
            "name": "wordpress",
            "state": {
              "running": {
                "startedAt": "2019-07-26T10:36:24Z"
              }
            },
            "lastState": {
              "terminated": {
                "exitCode": 137,
                "reason": "Error",
                "startedAt": "2019-07-25T12:38:59Z",
                "finishedAt": "2019-07-26T10:36:21Z",
                "containerID": "docker://fcf7968d80d1e5a2b4479ef7bf44f55097c56b60b8c4c96fa35b8de159f97008"
              }
            },
            "ready": true,
            "restartCount": 2,
            "image": "wordpress:4.8-apache",
            "imageID": "docker-pullable://wordpress@sha256:6216f64ab88fc51d311e38c7f69ca3f9aaba621492b4f1fa93ddf63093768845",
            "containerID": "docker://994457d8fdadb14c51655da08afb9ca716ac17d4e8ec8b448e1fce7b6976414a"
          }
        ],
        "qosClass": "BestEffort"
      }
    },
    {
      "metadata": {
        "name": "wordpress-mysql-684ddbbb55-zjs7b",
        "generateName": "wordpress-mysql-684ddbbb55-",
        "namespace": "wp1",
        "selfLink": "/api/v1/namespaces/wp1/pods/wordpress-mysql-684ddbbb55-zjs7b",
        "uid": "d37b56a7-3b32-46d0-9e95-849b45bf7c98",
        "resourceVersion": "2163328",
        "creationTimestamp": "2019-07-25T08:49:59Z",
        "labels": {
          "app": "wordpress",
          "pod-template-hash": "684ddbbb55",
          "tier": "mysql"
        },
        "ownerReferences": [
          {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "name": "wordpress-mysql-684ddbbb55",
            "uid": "438d61ec-60e9-4746-a892-15e352e1eb51",
            "controller": true,
            "blockOwnerDeletion": true
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "mysql-persistent-storage",
            "persistentVolumeClaim": {
              "claimName": "mysql-pvc-1"
            }
          },
          {
            "name": "default-token-7gh6x",
            "secret": {
              "secretName": "default-token-7gh6x",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "mysql",
            "image": "mysql:5.6",
            "ports": [
              {
                "name": "mysql",
                "containerPort": 3306,
                "protocol": "TCP"
              }
            ],
            "env": [
              {
                "name": "MYSQL_ROOT_PASSWORD",
                "value": "mysql@123"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "mysql-persistent-storage",
                "mountPath": "/var/lib/mysql"
              },
              {
                "name": "default-token-7gh6x",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-70-0-87-200.brbnca.spcsdns.net",
        "securityContext": {},
        "schedulerName": "stork",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:11Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:51:23Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:51:23Z"
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-07-25T08:50:11Z"
          }
        ],
        "hostIP": "70.0.87.200",
        "podIP": "192.168.6.137",
        "startTime": "2019-07-25T08:50:11Z",
        "containerStatuses": [
          {
            "name": "mysql",
            "state": {
              "running": {
                "startedAt": "2019-07-25T08:51:22Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "mysql:5.6",
            "imageID": "docker-pullable://mysql@sha256:104816d66fa44781b8229a431eec4987b5bc7ac6ba0011ee6d2738cce557a076",
            "containerID": "docker://133651f57e6f7a03a1367b4e47be567e3979f19e6dd5a2386f00de269284f615"
          }
        ],
        "qosClass": "BestEffort"
      }
    }
  ],
  "NodeMap": {
    "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2": {
      "id": "2bfcd2a4-d2aa-4708-beb7-2cba8f3f18f2",
      "cpu": 1.256281407035176,
      "mem_total": 8201998336,
      "mem_used": 1493413888,
      "mem_free": 6708584448,
      "status": 2,
      "mgmt_ip": "70.0.87.203",
      "data_ip": "70.0.87.203",
      "hostname": "ip-70-0-87-203.brbnca.spcsdns.net",
      "node_labels": {
        "Arch": "x86_64",
        "City": "Santa Clara",
        "Country": "United States",
        "Data IP": "70.0.87.203",
        "Docker Version": "18.03.0-ce",
        "ISP": "Telx",
        "ISP IP": "72.28.98.242",
        "Kernel Version": "3.10.0-862.3.2.el7.x86_64",
        "LAT": "3.73724E+01",
        "LNG": "-1.21974E+02",
        "Management IP": "70.0.87.203",
        "OS": "CentOS Linux 7 (Core)",
        "PX Version": "2.1.2.0-21409c7",
        "Region": "CA",
        "Timezone": "America/Los_Angeles",
        "Zip": "95051"
      },
      "scheduler_node_name": "ip-70-0-87-203.brbnca.spcsdns.net"
    },
    "4e056193-1f08-44f3-ac69-edf194937edd": {
      "id": "4e056193-1f08-44f3-ac69-edf194937edd",
      "cpu": 1.8820577164366374,
      "mem_total": 8201990144,
      "mem_used": 1450881024,
      "mem_free": 6751109120,
      "status": 2,
      "mgmt_ip": "70.0.87.233",
      "data_ip": "70.0.87.233",
      "hostname": "ip-70-0-87-233.brbnca.spcsdns.net",
      "node_labels": {
        "Arch": "x86_64",
        "City": "Santa Clara",
        "Country": "United States",
        "Data IP": "70.0.87.233",
        "Docker Version": "18.03.0-ce",
        "ISP": "Telx",
        "ISP IP": "72.28.98.242",
        "Kernel Version": "3.10.0-862.3.2.el7.x86_64",
        "LAT": "3.73724E+01",
        "LNG": "-1.21974E+02",
        "Management IP": "70.0.87.233",
        "OS": "CentOS Linux 7 (Core)",
        "PX Version": "2.1.2.0-21409c7",
        "Region": "CA",
        "Timezone": "America/Los_Angeles",
        "Zip": "95051"
      },
      "scheduler_node_name": "ip-70-0-87-233.brbnca.spcsdns.net"
    },
    "c5a1422f-5610-4d0a-9923-4b2e4cbb8693": {
      "id": "c5a1422f-5610-4d0a-9923-4b2e4cbb8693",
      "cpu": 1.1363636363636365,
      "mem_total": 8201998336,
      "mem_used": 1968275456,
      "mem_free": 6233722880,
      "status": 2,
      "mgmt_ip": "70.0.87.200",
      "data_ip": "70.0.87.200",
      "hostname": "ip-70-0-87-200.brbnca.spcsdns.net",
      "node_labels": {
        "Arch": "x86_64",
        "City": "Santa Clara",
        "Country": "United States",
        "Data IP": "70.0.87.200",
        "Docker Version": "18.03.0-ce",
        "ISP": "Telx",
        "ISP IP": "72.28.98.242",
        "Kernel Version": "3.10.0-862.3.2.el7.x86_64",
        "LAT": "3.73724E+01",
        "LNG": "-1.21974E+02",
        "Management IP": "70.0.87.200",
        "OS": "CentOS Linux 7 (Core)",
        "PX Version": "2.1.2.0-21409c7",
        "Region": "CA",
        "Timezone": "America/Los_Angeles",
        "Zip": "95051"
      },
      "scheduler_node_name": "ip-70-0-87-200.brbnca.spcsdns.net"
    }
  }
}
`
