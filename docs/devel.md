# Development
This page describes how to develop software for `pxc`.

# Overview
The goal of `pxc` is to provide a tool for users to use from their client machine like `kubectl` and `nomad`.

# Vendoring
Try to not vendor anything from `github.com/libopenstorage/openstorage`. `pxc` is an OpenStorage SDK application independent of openstorage packages. If you need a package from openstorage, check to see if you can just copy it to the `pkg/` directory.

# Commands
Please see [style](style.md) for more information on how to write commands

> NOTE: You may not need to add to pxc. Instead you may want to create your own component. See [Writing Components](components.md) for more information.

## Adding a command

* Create a new directory under `handler/`
* Add an import to the `handler/handler.go`
* Add the `get.go`, `create.go`, or any other other handlers in this new
  directory.
* Expose a `XXXAddCommand` function to let others attach commands to your new
  command.
* Attach your command to another by calling their `XXXAddCommand` function.
* See `handler/example` for example.

# Debugging

## Debugging using dlv on the command line

Example debugging `volume list`:

```
$ export KUBECONFIG=/your/kubeconfig
$ make install
$ PXC_KUBECTL_PLUGIN_MODE=true KUBECONFIG=/your/kubeconfig.conf dlv exec pxc -- volume list
```

Example of debugging `version`:

```
$ PXC_KUBECTL_PLUGIN_MODE=true dlv exec pxc -- version
Type 'help' for list of commands.
(dlv) b cmd/version.go:53
Breakpoint 1 set at 0x11c161e for github.com/portworx/pxc/cmd.versionExec() ./cmd/version.go:53
(dlv) c
> github.com/portworx/pxc/cmd.versionExec() ./cmd/version.go:53 (hits goroutine(1):1 total:1) (PC: 0x11c161e)
Warning: debugging optimized function
    48:         versionCmd.AddCommand(cmd)
    49: }
    50:
    51: func versionExec(cmd *cobra.Command, args []string) {
    52:         // Print client version
=>  53:         fmt.Printf("Client Version: %s\n"+
    54:                 "Client SDK Version: %s\n",
    55:                 PxVersion,
    56:                 fmt.Sprintf("%d.%d.%d",
    57:                         api.SdkVersion_Major,
    58:                         api.SdkVersion_Minor,
(dlv)
```

## Debugging using VSCode

First, check out this [blog](https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code).

Here is an example `launch.json` for pxc:

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "pxc version",
            "type": "go",
            "request": "launch",
            "mode": "exec",
            "program": "${workspaceRoot}/pxc",
            "env": {
                "PXC_KUBECTL_PLUGIN_MODE": "true",
                "KUBECONFIG": "${workspaceRoot}/kubeconfig.conf"
            },
            "args": [
                "version"
            ]
        },
        {
            "name": "pxc list-by-context",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceRoot}",
            "env": {
                "PXC_KUBECTL_PLUGIN_MODE": "true",
                "KUBECONFIG": "${workspaceRoot}/kubeconfig.conf"
            },
            "args": [
                "cluster",
                "list-by-context"
            ]
        }
    ]
}
```

## References

* [Help Screen Style](style.md)
* [Commands](commands.md)
