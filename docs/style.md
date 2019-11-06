# Help Style

Here is an example from `handler/volume/create.go`:

```go
	createVolumeCmd = &cobra.Command{
		Use:   "volume [NAME]",
		Short: "Create a volume in Portworx",
		Example: `
  # Create a volume called "myvolume" with size as 3GiB:
  pxc create volume myvolume --size=3

  # Create a volume called "myvolume" with size as 3GiB and replicas set to 3:
  pxc create volume myvolume --size=3 --replicas=3

  # Create a shared volume called "myvolume" with size as 3GiB:
  pxc create volume myvolume --size=3 --shared`,
  ...
```

* Two space left margin
* Comments start with a `#`
* Do not end comment in the examples with `:`
* Do not start the command with a `$ `, just use the command name
* Comments are above the line
* Start with the a verb in the present tense. Do not start the example with words like "This", or "To"
* Try to keep the comment to one sentence. If it is multiple sentences, then use a period to separate the lines.
* White space line between examples.
* The `Use:` must have `[NAME]` if the command takes the name as an argument.
* All commands are supposed to have `Use:`, `Short:`