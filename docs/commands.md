# Commands

## Remote Commands
Remote commands behave much the same way as the commands from `kubectl`.

### pxc get
`pxc get` is used to list objects from the server as done in `kubectl`. For example `pxc get volume` lists all the volumes in a table format.

#### `pxc get <object> <options...>`
`pxc get` with options can be used to output to json, yaml, or wide, where `wide` will show more columns.

#### `pxc get <object> <values...> <options...>`
If name of objects are passed to the command, then the list must only show those specific objects.

### pxc create
`pxc create` is used to create an object on Portworx. For example `pxc create clusterpair <options...>` creates a pair token at the server. Output must be a single line of just a `Created successfully` or something like `Created <object> <id/name> successfully>`.

We may have a future enhancement to do something like `pxc create -f <file.yml>` where the file may have many objects, but that will be done after 1.0.

### pxc delete
`pxc delete` is used to delete an object on Portworx. For example `pxc delete snapshot <name>` deletes a snapshot at the server. Output must be a single line of `Deleted <object> <name> successfully`.

### pxc describe
`pxc describe` is used to show information about an object on Portworx. For example `pxc describe volume <name>` will display volume information.

#### `pxc describe <object> <name/id> <options...>`
`pxc describe` with options can be used to output to json or yaml. If no option is provided, then the command shall print to the screen information about the object in simple format, not necessarily table format, as done in `kubectl describe...`

### pxc patch (update alias)
`pxc patch` is used to update an object in the server as done in `kubectl`. For example `pxc patch volume <name> --size=234` would resize the volume. Output must be a single line of `Updated <object> <name> successfully`.

## Local commands
Local commands act on the local system and may also interact with the server. Here are some example local commands:

* `pxc context`: Manages local context file
* `pxc cp <file> <volume:/path.../file>`: A possible `cp` command to move a local file to the volume