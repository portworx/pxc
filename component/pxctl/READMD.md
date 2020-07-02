# Component: pxctl

## Requirements

1. This requires bash, therefore only MacOS or Linux is supported. Windows users
can install pxc and this component under WSL2.
1. This component supports pxc only in kubectl plugin mode.

## Installation

Download the script above and save it under $HOME/.pxc/bin or anywhere in your
path. Once downloaded, make sure it is executable.

To check if pxc can find the component, execute the following command:

```
kubectl pxc component list
```

## Usage

You can then use it as a pxc compoment as follows:

```
kubectl pxc pxctl status
```

The script will first find a node, then execute _pxctl_ on that node. Not all
commands should be used, for example, _pxctl host attach_. Any commands that
affect the node should be executed on the node itself as normal.

If you want to run on a specific node, you can run the following:

```
kubectl pxc pxctl -n <node name> status
```

This will let the pxctl component know which node to execute the command.
