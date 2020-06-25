# Structure



### Code Layout

The tool should be structured so that it can execute any action against our
Tenant Clusters. There might be many cluster definitions for which we have many
actions each. The command line tool layout looks like shown below. In the
schematic example `cl001` is the cluster tested for conformity based on a
specific definition. All actions within that subcommand structure are
exclusively designed for this particular cluster layout, so it can be specific
to a single feature we want to test, e.g. single master clusters. In the example
below `ac001` would provide subcommands to execute or explain the action itself.

```
├── cmd
│   └── cl001
│       └── ac001
│           ├── execute
│           └── explain
└── pkg
    └── action
        └── ac001
```

```
$ awscnfm cl001 ac001 execute
$ awscnfm cl001 ac001 explain
$ awscnfm cl001 ac002 ...
```



### Execute Tests

Executing tests is silent by design. The command does not print anything and
silently exits with status code 0. In case the test failed an error is produced
and printed.

```
$ awscnfm cl001 ac001 execute
```



### Explain Tests

```
$ awscnfm cl001 ac001 explain
Check if the desired amount of Tenant Cluster master nodes are up and ready.

	* Fetch all G8sControlPlane CRs spec.replicas so that we know how many masters the Tenant Cluster is supposed to have.
	* Fetch the Tenant Cluster master nodes.
	* Compare the current and desired amount of master nodes.

```



### Tool Version

```
$ awscnfm version
```
