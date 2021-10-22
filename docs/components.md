# PXC Components

`pxc` supports "pluggable" components that enhance the workflow for the user. These components can be written in any language, even bash! The technology is based of [kubectl's plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) which you should read first.

`pxc` components written in golang have the benefit of using pxc libraries to manage the kubeconfig and connections to Portworx and Kubernetes.

Please see the examples and available components in [`components/`](https://github.com/portworx/pxc/tree/master/component) for more information.

## Environment Variables

`pxc` automatically passes environment variables to the component. These include all the current env variable of the local shell plus the following:

1. PXC_KUBECTL_PLUGIN_MODE: Boolean - Determine if pxc is running as a kubectl plugin
1. PXC_TOKEN: String - Portworx Token for Px Security, if any
1. PXC_PORTWORX_SERVICE_NAMESPACE: Service namespace of the SDK API
1. PXC_PORTWORX_SERVICE_NAME: Serivce name of the SDK API
1. PXC_PORTWORX_SERVICE_PORT: Service port of the SDK API

Also see: https://github.com/portworx/pxc/blob/master/pkg/util/utils.go#L32-L36


## Installation

A component can be installed on the `$PATH` or in `.pxc/bin` which defaults to `$HOME/.pxc/bin`. To determine if pxc can find your component, run the following command:

```
$ kubectl pxc component list
```
