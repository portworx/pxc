[![Build Status](https://travis-ci.org/portworx/px.svg?branch=master)](https://travis-ci.org/portworx/px)

# Overview
`px` is a client side application which communicates with Portworx, Kubernetes,
and other services to provide users with an integrated tool.

# Status: Pre-alpha
This tool is under heavy development.

# Downloads
Please refer to the [Releases](https://github.com/portworx/px/releases) page to
download the latest build.

# Usage
`px` is a tool that communicates with Portworx as well as Container
Orchestration systems like Kubernetes. To use this tool you must first create a
context with the appropriate information. `px` uses the context to connect to
the appropriate Portworx cluster to execute the requested command.

## Installing as a kubectl plugin
Install `px` binary anywhere in your PATH and name it `kubectl-px`. You will
then be able to run it like this:

```
$ kubectl px get nodes
$ kubectl px get pvc
```

## Creating a context
You can create a context using the following command:

```
$ px context create --name=mycluster --endpoint=<ip of cluster>:9020
```

See `px context create --help` for more information like setting the Kubeconfig.

> NOTE: The default gRPC SDK port for Portworx is 9020

### What if you don't have a Portworx cluster?
`px` uses the [OpenStorage SDK](https://libopenstorage.github.io) to communicate
with Portworx, therefore it is fully compatible with OpenStorage's
`mock-sdk-server`. If you do not have a Portworx cluster, you can run the
following to start the `mock-sdk-server`:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 openstorage/mock-sdk-server
$ px context create --name=mycluster --endpoint=localhost:9100
```

## Status
Now that `px` has been setup with a context, you can do the following commands:

```
$ px status
$ px get volume
$ px get volume -o wide
$ px get nodes
$ px get nodes -o wide
```

# Development
Please visit [Development](docs/devel.md)

