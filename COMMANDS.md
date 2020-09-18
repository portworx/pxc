# Commands

* Name of pxc or px as a plugin?
kubectl px
 -- Harsh to check name with Gou

* pxc config
  - kubectl pxc config set-cluster --portworx-namespace=asdf

* pxc describe pvc/cluster/volume <id>:
  pxc describe portworx-cluster ADVANCED
  pxc describe cluster --sdk-direct (HIDDEN)
  - cluster: Use the operator storageCluster object
  - pvc:
* pxc get cluster <id> -o yaml
* pxc get clusters
....

# Development
* Have operator write it's handler in their tree if they want

# AIs
* Operator documentation should say that the StorageCluster/Node etc should
  be accessed only by those with permission.
* Add security support to the Operator
* No advanced mode. just keep them hidden
* Proxy Node specific calls, like "place node X to maintenance mode"

