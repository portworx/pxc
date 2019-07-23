# Development
This page describes how to develop software for `px`.

# Overview
The goal of `px` is to provide a tool for users to use from their client machine like `kubectl` and `nomad`. `px` is based on the `kubectl` style of commands lowering the rampup to learn this tool.

# Root commands
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
`px create` is used to create an object on Portworx. For example `px create clusterpair` creates a pair token at the server. Output must be a single line of just a "Created successfully" or printing the id of the object created.

#### `px create <object> <options...>`

#### `px create <object> <values...> <options...>`


