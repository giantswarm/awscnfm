# Using awscnfm

This article assumes that all you want to do is using existing cluster scopes
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

You need access to and credentials for a (test)
installation's Control Plane Kubernetes API.

`awscnfm` assumes you have a kubeconfig file somewhere with the context set to
that control plane. It looks at the `AWSCNFM_CONTROLPLANE_KUBECONFIG`
environment variable to find out about the path to that file.

Here is an example:

```bash
export AWSCNFM_CONTROLPLANE_KUBECONFIG=~/.kube/config
```

## 3. Select a cluster scope

Running the `awscnfm` CLI without any arguments gives you an overview of all the
subcommands. Commands in the format `cl***` (e. g. `cl001`) are cluster scopes.

To learn more about each cluster scope available, check the help usage of the
first action in the scope.

For `cl001` that would be:

```nohighlight
awscnfm cl001 ac000 -h
```

The output will look similar to this

```nohighlight
Execute the conformance test plan of this cluster scope. Actions below are
executed in order. A tenant cluster is conform if the plan executes without
errors. Plan execution might take up to 2h19m10s.

ACTION  RETRY  WAIT     COMMENT
ac001   2s     10s      create cluster CRs
ac002   3m0s   24m0s    check cluster access
ac003   2s     10s      check master count
ac004   2s     10s      check worker count
ac005   2s     10s      create node pool
ac004   3m0s   24m0s    check worker count
ac006   2s     10s      check master host network
ac007   2s     10s      check worker host network
ac008   2s     10s      delete cluster CRs
ac009   9m0s   1h30m0s  check CRs deleted
```

## 4. Execute a plan or an action

You can choose to execute either the entire test plan for a cluster scope, or a single action.

### 4a. Executing the entire plan

Running the entire plan, that is all actions defined within the cluster scope, should help you

- create a tenant cluster according to the spec of the cluster scope
- run all checks
- delete the cluster in the end

For example, with the cluster scope `cl001`, executing the entire plan would require this command:

```nohighlight
awscnfm cl001 ac000
```

Note that the action name argument `ac000` represents the entire plan.

It may take some time until you will see some output.

### 4b. Executing a single action

Instead of setting the action argument to `ac000` for the entire test plan, you can use any other available action identifier with the execute command. Example:

```nohighlight
awscnfm cl001 ac007
```

However, since the cluster creation step will not be executes now, you have to make sure that

- there is a **cluster** to use, with specifications matching the cluster scope and
- `awscnfm` knows the cluster ID

There are two ways to do this:

#### 1. Scope configuration file

After launching the above command, `awscnfm` will create a new cluster. The ID of this cluster will be persisted in the file

```nohighlight
~/.config/awscnfm/cl001.yaml
```

You can modify this file to change the cluster ID.

#### 2. Environment variable

The cluster ID setting in the configuration file above can be overwritten via the environment variable

```nohighlight
AWSCNFM_TENANTCLUSTER
```

## 5. Cleaning up

Make sure to clean up tenant clusters after testing, especially after executing single actions.

One way to do this is to execute the action provided for that purpose.
