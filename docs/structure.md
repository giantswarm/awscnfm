# Repository structure

The tool should be structured so that it can execute any action against our
Tenant Clusters. There might be many cluster definitions for which we have many
actions each. The command line tool layout looks like shown below. In the
schematic example `cl001` is the cluster scope for testing conformity based on a
specific tenant cluster definition. All actions within that subcommand structure
are exclusively designed for this particular cluster scope, so it can be
specific to a single feature we want to test, e.g. single master clusters. In
the example below `ac001` would provide subcommands to execute or explain the
action itself. Note that all actions defined within a certain cluster scope are
meant to be open and generic. Any business logic can be implemented. The only
assumption we make for the first action `ac001` is that it is meant to create
the tenant cluster for its own particular cluster scope. This implies to not
create a Kubernetes client for the tenant cluster when executing the first
action of a given cluster scope.

```nohighlight
├── cmd
│   └── cl001
│       └── ac001
│           ├── execute
│           └── explain
└── pkg
    └── action
        └── ac001
```

The according subcommand hierarchy:

```nohighlight
awscnfm cl001 ac001 execute
awscnfm cl001 ac001 explain
awscnfm cl001 ac002 ...
```
