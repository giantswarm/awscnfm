# Environment



### Control Plane Access

In order to create a rest config for the Control Plane we want to work with, the
following environment variable must be set. Note that this assumes you have a
working connection setup and active to the Control Plane defined in your kube
config.

```
export AWSCNFM_KUBECONFIG=~/.kube/config
```



### Tenant Cluster Access

The first action of a cluster scope is meant to create the CRs necessary to
bootstrap a new tenant cluster. The resulting tenant cluster ID is written to a
config file in your local file system.

```
$ cat ~/.config/awscnfm/cl001.yaml
Cluster: al9qy
```

In order to create a rest config for the Tenant Cluster we want to work with,
further sub commands and actions will look for the generated tenant cluster ID
in the config file described above. Note that this mechanism uses the
`cluster-operator` specific API certificates found in the Control Plane. It is
also possible to make use of the following environment variable in order to
overwrite the used tenant cluster ID.

```
export AWSCNFM_TENANTCLUSTER=al9qy
```
