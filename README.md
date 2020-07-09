[![Build Status](https://travis-ci.com/portworx/pxc.svg?token=koUsyDmAMgMD5TViiacc&branch=master)](https://travis-ci.com/portworx/pxc)

# Overview
`pxc` is a client side application which communicates with Portworx, Kubernetes,
and other services to provide users with an integrated tool. It can be used as
a stand alone program or as a kubectl plugin.

`pxc` also support pluggable runtime components. See [`component/`](component)
for more information.

# Downloads
Please refer to the [Releases](https://github.com/portworx/pxc/releases) page to
download the latest build.

# Documentation

Please see [documentation](docs/usage/pxc.md)

# Usage

## Kubectl Plugin
Install `kubectl-pxc` binary anywhere in your PATH. You will
then be able to run it like this:

```
$ kubectl pxc cluster describe
$ kubectl pxc node list
$ kubectl pxc pvc list
```

## Standalone

### On a Portworx node
When not configured, `pxc` defaults to using the local `127.0.0.1:9020` port on the host.
You will not need to do any configuration if you install pxc on a Portworx node where it
has setup the SDK gRPC server on port 9020.

### From a client machine

Normally, you would run `pxc` from a client machine. Once you download the `pxc` binary,
you will need to configure it. Here is an example:

```
 ./pxc config cluster set --name=clusterone endpoint=1.1.1.1:9020
 ./pxc config context set --cluster=clusterone --name=contextone
 ./pxc config context use --name=contextone
 ./pxc cluster describe
```

# Development
Please visit [Development](docs/devel.md)

