# Development
This page describes how to develop software for `px`.

# Overview
The goal of `px` is to provide a tool for users to use from their client machine like `kubectl` and `nomad`. `px` is based on the `kubectl` style of commands lowering the rampup to learn this tool.

# Commands
There are two style of commands: Local and remote commands.

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
`px delete` is used to delete an object on Portworx. For example `px delete snapshot <name>` deletes a snapshot at the server. Output must be a single line of just a `Deleted <object> <name> successfully`.

### px describe
`px describe` is used to show information about an object on Portworx. For example `px describe volume <name>` will display volume information.

#### `px describe <object> <name/id> <options...>`
`px describe` with options can be used to output to json or yaml. If no option is provided, then the command shall print to the screen information about the object in simple format, not necessarily table format, as done in `kubectl describe...`


