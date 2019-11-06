[![Build Status](https://travis-ci.com/portworx/pxc.svg?token=koUsyDmAMgMD5TViiacc&branch=master)](https://travis-ci.com/portworx/pxc)

# Overview
`pxc` is a client side application which communicates with Portworx, Kubernetes,
and other services to provide users with an integrated tool.

# Downloads
Please refer to the [Releases](https://github.com/portworx/pxc/releases) page to
download the latest build.

# Usage
`pxc` is a tool that communicates with Portworx as well as Container
Orchestration systems like Kubernetes. It can be run standalone or as a plugin
to kubectl.

## Kubectl Plugin
Install `pxc` binary anywhere in your PATH and name it `kubectl-pxc`. You will
then be able to run it like this:

```
$ kubectl pxc get nodes
$ kubectl pxc get pvc
$ kubectl pxc get pvc --kubeconfig=/path/to/kubeconfig.conf
```

If you are running Portworx installed on a Kubernetes Cloud like GKE, EKS, etc,
please use this model.

`pxc` will automatically discover how to communicate with Portworx. No need for
any prior setup.

## Standalone 

You must first create a context with the appropriate information. `pxc` uses the context to connect to
the appropriate Portworx cluster to execute the requested command.

### Creating a context
You can create a context using the following command:

```
$ pxc context create mycluster --endpoint=<ip of cluster>:9020 --kubeconfig=/path/to/kubeconfig
```

> NOTE: The default gRPC SDK port for Portworx is 9020

#### What if you don't have a Portworx cluster?
`pxc` uses the [OpenStorage SDK](https://libopenstorage.github.io) to communicate
with Portworx, therefore it is fully compatible with OpenStorage's
`mock-sdk-server`. If you do not have a Portworx cluster, you can run the
following to start the `mock-sdk-server`:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 openstorage/mock-sdk-server
$ pxc context create --name=mycluster --endpoint=localhost:9100
```

### pxc sample commands
Now that `pxc` has been setup with a context, you can do the following commands:

```
$ pxc describe cluster
$ pxc get volume
$ pxc get volume -o wide
$ pxc get nodes
$ pxc get nodes -o wide
```

# Development
Please visit [Development](docs/devel.md)

