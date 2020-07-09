## pxc utilities token-generate

Generate a Portworx token

### Synopsis

Generate a Portworx token

```
pxc utilities token-generate [flags]
```

### Examples

```

  # Login to portworx using a secret in Kubernetes
  pxc utilities token-generate pxc util token-generate \
	--token-email=example.user@example.com \
	--token-name="Example User" \
	--token-roles=system.user \
	--token-groups=exampleGroup \
	--token-duration=7d \
	--token-subject="exampleCompany/example.user@example.com" \
	--shared-secret=mysecret
```

### Options

```
      --ecdsa-private-keyfile string   ECDSA Private file to sign token
  -h, --help                           help for token-generate
      --rsa-private-keyfile string     RSA Private file to sign token
      --shared-secret string           Shared secret to sign token
      --token-duration string          Duration of time where the token will be valid. Postfix the duration by using s for seconds, m for minutes, h for hours, d for days, and y for years. (default "1d")
      --token-email string             Unique ID of this account
      --token-groups string            Comma separated list of groups which the token will be part of
      --token-issuer string            Issuer name of token. Do not use https:// in the issuer since it could indicate that this is an OpenID Connect issuer. (default "portworx.com")
      --token-name string              Account name
      --token-roles string             Comma separated list of roles applied to this token
      --token-subject string           Unique ID of this account
```

### Options inherited from parent commands

```
      --pxc.config string       Config file (default is $HOME/.pxc/config.yml)
      --pxc.config-dir string   Config directory (default "/home/lpabon/.pxc")
      --pxc.context string      Force context name for the command
      --pxc.token string        Portworx authentication token
      --pxc.v int32             [0-3] Log level verbosity
```

### SEE ALSO

* [pxc utilities](pxc_utilities.md)	 - pxc utility commands

###### Auto generated by spf13/cobra on 2-Jul-2020