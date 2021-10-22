## pxc login

Set authentication information for Portworx cluster

### Synopsis

Saves your Portworx authentication information for the current
user in the kubeconfig file for future access of the Portworx system.

```
pxc login [flags]
```

### Examples

```

  # Login to portworx using a secret in Kubernetes
  pxc login --k8s-secret-name=abc --k8s-secret-namespace=ns

  # Login to portworx using a specified token
  pxc login --auth-token=eyJh...sb30ro
```

### Options

```
      --auth-token string             Auth token if any (optional)
  -h, --help                          help for login
      --k8s-secret-name string        Kubernetes secret name with the auth token
      --k8s-secret-namespace string   Kubernetes namespace containing the secret with the auth token
```

### Options inherited from parent commands

```
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default cache directory (default "/home2/lpabon/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
      --pxc.config string              Config file (default is $HOME/.pxc/config.yml)
      --pxc.config-dir string          Config directory (default "/home2/lpabon/.pxc")
      --pxc.secret-name string         Kubernetes secret name containing authentication token
      --pxc.secret-namespace string    Kubernetes namespace where secret contains token
      --pxc.token string               Portworx authentication token
      --pxc.v int32                    [0-3] Log level verbosity
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
```

### SEE ALSO

* [pxc](pxc.md)	 - Portworx client

###### Auto generated by spf13/cobra on 29-Sep-2021