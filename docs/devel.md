# Development
This page describes how to develop software for `pxc`.

# Overview
The goal of `pxc` is to provide a tool for users to use from their client machine like `kubectl` and `nomad`. `pxc` is based on the `kubectl` style of commands lowering the rampup to learn this tool.

Keep in mind that `pxc` will also be used as a plugin to `kubectl`. Therefore commands like `kubectl pxc get volumes` and `kubectl pxc get pvc` must be clear to the user.

# Vendoring
Try to not vendor anything from `github.com/libopenstorage/openstorage`. `pxc` is an OpenStorage SDK application independent of openstorage packages. If you need a package from openstorage, check to see if you can just copy it to the `pkg/` directory.

# Commands
There are two style of commands: Local and remote commands.

## Adding a command

* Create a new directory under `handler/`
* Add an import to the `handler/handler.go`
* Add the `get.go`, `create.go`, or any other other handlers in this new
  directory.
* Expose a `XXXAddCommand` function to let others attach commands to your new
  command.
* Attach your command to another by calling their `XXXAddCommand` function.
* See `handler/example` for example.

## References

* [Help Screen Style](style.md)
* [Commands](commands.md)
