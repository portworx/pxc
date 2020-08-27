# PXC Components

`pxc` supports "pluggable" components that enhance the workflow for the user. These components can be written in any language, even bash! The technology is based of [kubectl's plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) which you should read first.

`pxc` components written in golang have the benefit of using pxc libraries to manage the kubeconfig and connections to Portworx and Kubernetes.

Please see the examples and available components in [`components/`](https://github.com/portworx/pxc/tree/master/component) for more information.

## Installation

A component can be installed on the `$PATH` or in `.pxc/bin` which defaults to `$HOME/.pxc/bin`. To determine if pxc can find your component, run the following command:

```
$ kubectl pxc component list
```
