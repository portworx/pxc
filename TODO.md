# Doc

* Two space left margin
* Comments start with a `#`
* Do not end comment in the examples with `:`
* Do not start the command with a `$ `, just use the command name
* Commaents are above the line
* Start with the a verb in the present tense. Do not start the example with words like "This", or "To". The comment can be the what the command does in English.
* Try to keep the comment to one sentence. If it is multiple sentences, then use a period to separate the lines.
* White space line between examples.
* The `Use:` must have `[NAME]` if the command takes the name as an argument.
* The `Use:` must have shown any non-subcommand arguments as `[VALUE]` all in caps.
* All commands are supposed to have `Use:`, `Short:`
* All handler commands should expose a function to add external commands to it.

# TODO

* Convert STATUS_OK to something
* move cluster_test package to cluster
* move context_test package
* move volume_test package
* move volumeclone_test
* move volumesnapshot_test
* Do we need context unset?