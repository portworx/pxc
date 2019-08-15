# Development
This page describes how to develop software for `px`.

# Overview
The goal of `px` is to provide a tool for users to use from their client machine like `kubectl` and `nomad`. `px` is based on the `kubectl` style of commands lowering the rampup to learn this tool.

Keep in mind that `px` will also be used as a plugin to `kubectl`. Therefore commands like `kubectl px get volumes` and `kubectl px get pvc` must be clear to the user.

# Vendoring
Try to not vendor anything from `github.com/libopenstorage/openstorage`. `px` is an OpenStorage SDK application independent of openstorage packages. If you need a package from openstorage, check to see if you can just copy it to the `pkg/` directory.

# Commands
There are two style of commands: Local and remote commands.

## Adding a command

* Install [cobra](https://github.com/spf13/cobra#installing)
* Type: `cobra add <verb>` to add a root command

* Adding a subcommand: `cobra add <verbName> -p <parent>Cmd` for example: `cobra add getPv -p getCmd`

## Remote Commands
Remote commands behave much the same way as the commands from `kubectl`.

### px get
`px get` is used to list objects from the server as done in `kubectl`. For example `px get volume` lists all the volumes in a table format.

#### `px get <object> <options...>`
`px get` with options can be used to output to json, yaml, or wide, where `wide` will show more columns.

#### `px get <object> <values...> <options...>`
If name of objects are passed to the command, then the list must only show those specific objects.

### px create
`px create` is used to create an object on Portworx. For example `px create clusterpair <options...>` creates a pair token at the server. Output must be a single line of just a `Created successfully` or something like `Created <object> <id/name> successfully>`.

We may have a future enhancement to do something like `px create -f <file.yml>` where the file may have many objects, but that will be done after 1.0.

### px delete
`px delete` is used to delete an object on Portworx. For example `px delete snapshot <name>` deletes a snapshot at the server. Output must be a single line of `Deleted <object> <name> successfully`.

### px describe
`px describe` is used to show information about an object on Portworx. For example `px describe volume <name>` will display volume information.

#### `px describe <object> <name/id> <options...>`
`px describe` with options can be used to output to json or yaml. If no option is provided, then the command shall print to the screen information about the object in simple format, not necessarily table format, as done in `kubectl describe...`

### px patch (update alias)
`px patch` is used to update an object in the server as done in `kubectl`. For example `px patch volume <name> --size=234` would resize the volume. Output must be a single line of `Updated <object> <name> successfully`.

## Local commands
Local commands act on the local system and may also interact with the server. Here are some example local commands:

* `px context`: Manages local context file
* `px cp <file> <volume:/path.../file>`: A possible `cp` command to move a local file to the volume

# Extending commands

* Create a new directory under `handler/`
* Add an import to the `handler/handler.go`
* Add the `get.go`, `create.go`, or any other other handlers in this new
  directory.
* Expose a `XXXAddCommand` function to let others attach commands to your new
  command.
* Attach your command to another by calling their `XXXAddCommand` function.

