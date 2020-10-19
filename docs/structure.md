# Repository Structure

The `awscnfm` tool should be structured so that any action can be executed
separately against our Tenant Clusters. There might be many cluster definitions
for which we have many actions each. The project structure of the command line
tool looks like shown below. In the schematic example `cl001` is a cluster
scope. A cluster scope is for testing conformity based on a specific Tenant
Cluster definition. All actions within that subcommand structure are exclusively
designed for this particular cluster scope, so it can be specific to a single
feature we want to test, e.g. single master clusters. In the example below
`ac001` would provide business logic that can be executed in isolation. Note
that all actions defined within a certain cluster scope are meant to be generic.
Any business logic can be implemented. One assumption we make for the action
`ac000` is that it implements an execution plan. Another assumption we make is
that action `ac001` creates the Tenant Cluster for its own particular cluster
scope. This implies e.g. to not create a Kubernetes client for the Tenant
Cluster when executing action `ac001` of a given cluster scope.

The package hierarchy:

```nohighlight
├── cmd
│   └── cl001
│       └── ac001
└── pkg
    └── action
        └── ac001
```

The according subcommand hierarchy:

```nohighlight
awscnfm cl001 ac001
awscnfm cl001 ac002
awscnfm cl001 ac003
```
