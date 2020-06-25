# Environment

In order to create a rest config for the Tenant Cluster we want to work with,
the following environment variable must be set. Note that this uses the
`cluster-operator` specific API certificates found in the Control Plane.

```
export AWSCNFM_TENANTCLUSTER=al9qy
```

In order to create a rest config for the Control Plane we want to work with, the
following environment variable must be set. Note that this assumes you have a
working connection setup and active to the Control Plane defined in your kube
config.

```
export AWSCNFM_KUBECONFIG=~/.kube/config
```
