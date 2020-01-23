## pxc get alerts

Get information about Portworx alerts

### Synopsis

Get information about Portworx alerts

```
pxc get alerts [flags]
```

### Examples

```

  # Get portworx related alerts
  pxc get alerts

  # Fetch alerts based on particular alert id. Fetch all alerts based on "VolumeCreateSuccess" id
  pxc get alerts --id "VolumeCreateSuccess"

  # Fetch alerts between a time window
  pxctl alerts show --start-time "2019-09-19T09:40:26.371Z" --end-time "2019-09-19T09:43:59.371Z"

  # Fetch alerts with min severity level
  pxc get alerts --severity "alarm"

  # Fetch alerts based on resource type. Here we fetch all "volume" related alerts
  pxc get alerts -t "volume"

  # Fetch alerts based on resource id. Here we fetch alerts related to "cluster"
  pxc get alerts --id "1f95a5e7-6a38-41f9-9cb2-8bb4f8ab72c5"
```

### Options

```
  -e, --end-time string     end time span (RFC 3339)
  -h, --help                help for alerts
  -i, --id string           Alert id 
  -o, --output string       Output in yaml|json|wide
  -v, --severity string     Min severity value (Valid Values: [notify warn warning alarm]) (default "notify") (default "notify")
  -a, --start-time string   start time span (RFC 3339)
  -t, --type string         alert type (Valid Values: [volume node cluster drive all]) (default "all")
```

### Options inherited from parent commands

```
      --config string    Config file (default is $HOME/.pxc/config.yml)
      --context string   Force context name for the command
      --v int32          [0-4] Log level verbosity
```

### SEE ALSO

* [pxc get](pxc_get.md)	 - Get information from Portworx

###### Auto generated by spf13/cobra on 6-Nov-2019