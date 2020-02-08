# Portworx Command Client

## Command Structure

pxc commands have the following structure:

gcloud/pxctl style:

```
pxc [COMPONENT] [GROUP(s)] [COMMAND] [OPTIONS]
```

Kubernetes style:

```
pxc [COMPONENT] [VERB] [OBJ] [OPTIONS]
```

Only Kubernetes specific components are encorougaed to use the Kubernetes style.
All other components should use the gcloud/pxctl style.

## gcloud style

### Components

A component is an optional service provided by Portworx. Like central,
backup, DR, etc. A component can have `[GROUP] [COMMAND] [OPTIONS]`.

This method is used by most CLIs and it may benefit supporting plugins
in future releases.

#### Examples

```
# central as the component
pxc central login

# operator as the component
pxc operator version

# pxc as a component
pxc version
```

### Groups

Some components may have sub-components or groups as a way of s a grouping
commands for a certain task. Groups can have groups, commands/verbs, or options.

#### Examples

```
# cluster is the group
pxc central cluster list
```

### Commands

Commands or verbs take options as flags and they are the actions on an object.

#### Examples

```
# login is the command
pxc central login

# status is the command
pxc operator status -o=yaml

# set-cluster is the command
pxc config set-cluster --endpoint=...
```

### Kubernetes Style

CLI capabilities specific to Kuberntes are encouraged to use the Kubernetes
style.

#### Examples

```
# stork
pxc stork create migration mymigration
```

