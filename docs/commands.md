# Portworx Command Client

## Command Structure

pxc commands have the following structure which is based on the gcloud/pxctl style:

```
pxc [COMPONENT] [GROUP(s)] [COMMAND] [OPTIONS]
```

### Components

A component is an optional service provided by Portworx. Like central,
backup, DR, etc. A component can have `[GROUP] [COMMAND] [OPTIONS]`.

This method is used by most CLIs and it may benefit supporting plugins
in future releases.

For a component example see
[pxc-component-example](https://github.com/portworx/pxc-component-example)

#### Example component commands

```
# central as the component
pxc central login

# operator as the component
pxc operator version
```

### Groups

Some components may have sub-components or groups as a way of grouping
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


## Kubernetes style grandfathered

Some current applications may already be using the Kubernetes style:

```
pxc [COMPONENT] [VERB] [OBJ] [OPTIONS]
```

These applications may still need to use this model to avoid confusing current
users.
