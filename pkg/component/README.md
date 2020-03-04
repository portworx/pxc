[![Build Status](https://travis-ci.com/portworx/pxc-component-example.svg?token=L4phF4S8kCmz2B53gLyB&branch=master)](https://travis-ci.com/portworx/pxc-component-example)

# Example pxc component
A pxc component is a plugin for pxc which utilizes the pxc libraries to connect
to Portworx or Kubernetes.

# Usage

## Installation

1. Type: `make clean all`
1. Install `pxc-cm` either in your path or in `$HOME/.pxc/bin`

## Usage

Now you can use it as a kubectl-pxc plugin as follows:

```
$ export KUBECONFIG=k8s1.conf:k8s2.conf

$ kubectl pxc cm show --context=k8s2-admin@k8s2
Cluster ID: px-cluster-ad340bb9-be1e-4acc-94e6-c2eff0f7a47f
Cluster UUID: 142d6f2d-0f43-400c-8b23-f14294907bd3
Cluster Status: Ready
Version: 2.3.2.0-c6ebdd7
 build: c6ebdd7ab5801b2d0086179610e5848236e1b95c
SDK Version 0.42.22

$ kubectl pxc cm show
Cluster ID: px-cluster-43e75a0d-39d5-4904-ae11-b6f00c51f261
Cluster UUID: f7076dba-b700-472f-945e-9c6b310de76b
Cluster Status: Ready
Version: 2.2.0.4-c964260
 build: c9642608c55c9a5dc109296822126afc35a0eb9e
SDK Version 0.42.17

$ kubectl pxc cm show -o json
```

As you can see above, the component `cm` (cluster manager) uses the pxc libraries to
automatically connect to the appropriate Portworx system according to the context
in kubeconfig.


