[![Build Status](https://travis-ci.com/portworx/pxc.svg?token=koUsyDmAMgMD5TViiacc&branch=master)](https://travis-ci.com/portworx/pxc)

# Overview
`pxc` is a client side application which communicates with Portworx, Kubernetes,
and other services to provide users with an integrated tool.

The first release of pxc is focused on being a kubectl plugin.

# Downloads
Please refer to the [Releases](https://github.com/portworx/pxc/releases) page to
download the latest build.

# Documentation

Please see [documentation](docs/usage/pxc.md)

# Usage
`pxc` is a command line client that communicates with Portworx as well as Container
Orchestration systems like Kubernetes. It can be run standalone or as a plugin
to kubectl.

## Kubectl Plugin
Install `kubectl-pxc` binary anywhere in your PATH. You will
then be able to run it like this:

```
$ kubectl pxc get nodes
$ kubectl pxc get pvc
$ kubectl pxc get pvc --kubeconfig=/path/to/kubeconfig.conf
```

# Development
Please visit [Development](docs/devel.md)

