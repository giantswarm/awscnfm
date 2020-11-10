# Structure

The `awscnfm` tool should be structured so that any action can be executed
separately against our Tenant Clusters. There might be many cluster definitions
for which we have many actions each. A test plan is for testing conformity based
on a specific Tenant Cluster definition. Any business logic can be implemented
in dedicated actions. Actions can be reused across test plans. The project
structure of the command line tool looks like shown below.

```
cmd
├── action
│   ├── create
│   │   ├── cluster
│   │   │   ├── customnetworkpool
│   │   │   ├── defaultcontrolplane
│   │   │   └── singlecontrolplane
│   │   ├── kiam
│   │   │   └── awsapicall
│   │   └── nodepool
│   │       └── defaultdataplane
│   ├── delete
│   │   ├── cluster
│   │   ├── kiam
│   │   │   └── awsapicall
│   │   └── networkpool
│   ├── update
│   │   └── cluster
│   │       ├── hacontrolplane
│   │       ├── major
│   │       ├── minor
│   │       └── patch
│   └── verify
│       ├── apps
│       │   └── installed
│       ├── cluster
│       │   ├── created
│       │   ├── deleted
│       │   ├── ha
│       │   ├── networkpool
│       │   └── updated
│       ├── kiam
│       │   ├── awsapicall
│       │   └── podandsecret
│       ├── master
│       │   ├── hostnetworkpod
│       │   └── ready
│       └── worker
│           ├── hostnetworkpod
│           └── ready
└── plan
    ├── pl001
    ├── pl002
    ├── pl003
    ├── pl004
    ├── pl005
    └── pl006
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
