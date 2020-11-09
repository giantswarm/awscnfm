# Repository Structure

The `awscnfm` tool should be structured so that any action can be executed
separately against our Tenant Clusters. There might be many cluster definitions
for which we have many actions each. The project structure of the command line
tool looks like shown below. In the schematic example `cl001` is a cluster
scope. A test plan is for testing conformity based on a specific Tenant
Cluster definition. All actions within that subcommand structure are exclusively
designed for this particular test plan, so it can be specific to a single
feature we want to test, e.g. single master clusters. In the example below
`ac001` would provide business logic that can be executed in isolation. Note
that all actions defined within a certain test plan are meant to be generic.
Any business logic can be implemented. One assumption we make for the action
`ac000` is that it implements an execution plan. Another assumption we make is
that action `ac001` creates the Tenant Cluster for its own particular cluster
scope. This implies e.g. to not create a Kubernetes client for the Tenant
Cluster when executing action `ac001` of a given test plan.

An example package structure looks like below.

```nohighlight
├── cmd
│   ├── action
│   │   ├── create
│   │   │   └── cluster
│   │   │       ├── default
│   │   │       └── networkpool
│   │   ├── delete
│   │   │   └── cluster
│   │   └── verify
│   │       └── cluster
│   │           ├── created
│   │           └── deleted
│   └── plan
│       ├── pl001
│       ├── pl002
│       └── pl003
└── pkg
    ├── action
    │   ├── create
    │   │   └── cluster
    │   │       ├── default
    │   │       └── networkpool
    │   ├── delete
    │   │   └── cluster
    │   └── verify
    │       └── cluster
    │           ├── created
    │           └── deleted
    └── plan
        ├── pl001
        ├── pl002
        └── pl003
```

An example command structure looks like below.

```nohighlight
awscnfm action create cluster default
awscnfm action create cluster networkpool
awscnfm action verify cluster created
awscnfm action verify cluster deleted
awscnfm plan pl001
awscnfm plan pl002
awscnfm plan pl003
```
