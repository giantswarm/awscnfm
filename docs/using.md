# Using awscnfm

This article assumes that all you want to do is using existing test plans
and actions.



## 1. Decide on the version to use

`awscnfm` has a versioning scheme aligned with the Giant Swarm tenant cluster
releases for AWS. Both the major and the minor version must match. Assuming
`v11.5.1` is the latest patch release of `v11.5.x`, you should use the latest
`awsncfm` release in minor version `v11.5.x`. If you want to specifically test
`v11.5.1` even though there does `v11.5.3` exist, you need to specify the
release version using the environment variable `AWSCNFM_RELEASEVERSION`. That
way you can then also test any test releases like `v100.0.0-xh3b4sd`.

In order to find a particular release, either download and unpack the
[release](https://github.com/giantswarm/awscnfm/releases) or clone the source
repository, check out the relevant Git Tag, and `go build` the binary yourself.



## 2. Get control plane access

You need access to and credentials for a (test) installation's Control Plane
Kubernetes API. `awscnfm` assumes you have a kubeconfig file somewhere with the
context set to that control plane. It looks at the environment variable
`AWSCNFM_CONTROLPLANE_KUBECONFIG` to find out about the path to that file. The
default is already your local kube config.

```
AWSCNFM_CONTROLPLANE_KUBECONFIG=~/.kube/config
```



## 3. Execute a test plan

Running the `awscnfm` CLI without any arguments gives you an overview of all the
subcommands. Autocompletion helps you discovering the command line tool. The
help usage of a test plan command provides more information about the executed
steps.

```
$ awscnfm plan pl003 -h
Test plan pl003 launches a basic Tenant Cluster in the previous patch release
and upgrades the Tenant Cluster to the latest patch release once it is up. Plan
execution might take up to 4h30m50s.

ACTION                              RETRY  WAIT
create/cluster/defaultcontrolplane  2s     10s
verify/cluster/created              3m0s   30m0s
verify/master/ready                 2s     10s
create/nodepool/defaultdataplane    2s     10s
verify/worker/ready                 3m0s   30m0s
update/cluster/minor                2s     10s
verify/cluster/updated              10m0s  2h0m0s
delete/cluster                      2s     10s
verify/cluster/deleted              9m0s   1h30m0s

Usage:
  awscnfm plan pl003 [flags]

Flags:
  -h, --help   help for pl003
```



## 4. Executing a single action

You can choose to execute a single action. The help usage of an action command
provides more information about the executed steps.

```
$ awscnfm action verify cluster updated -h
Check if the Tenant Cluster got successfully upgraded. Note that this
particular action is not meant to be reliably used for other purposes than
for the plan exection. Executing this action against a Tenant Cluster that
got already upgraded may lead to wrong results in case you want to assert an
additional Tenant Cluster upgrade.

    * Fetch the AWSCluster CR.
    * Check if the latest cluster status condition is "Updated".
    * Return an error if we see other cluster status conditions than "Updated".

Usage:
  awscnfm action verify cluster updated [flags]

Flags:
  -h, --help                    help for updated
  -c, --tenant-cluster string   Tenant Cluster ID to use for this particular action.
```



## 5. Cleaning up

Make sure to clean up tenant clusters after testing, especially after executing
single actions. One way to do this is to execute the action provided for that
purpose.

```
$ awscnfm action delete cluster -h
Delete all Tenant Cluster CRs on the Control Plane by triggering the
deletion of the Cluster CR. This should ensure the following.

    * Trigger deletion to all other CRs associated with the tenant cluster.
    * Execute cleanup logic in all involved operators.
    * Remove all cloud provider resources.
    * Remove all CRs associated with the tenant cluster.

Usage:
  awscnfm action delete cluster [flags]

Flags:
  -h, --help                    help for cluster
  -c, --tenant-cluster string   Tenant Cluster ID to use for this particular action.
```
