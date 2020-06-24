[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](http://godoc.org/github.com/giantswarm/awscnfm)

# awscnfm

Conformance test utility for AWS.



### Setup

As of now you need to start your own Tenant Cluster so that you can execute the
implemented conformance tests against it. Cluster creation can be another action
supported in the test suite at a later point in time.



### Structure

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
│   ├── cl001
│       └── ac001
│           ├── execute
│           └── explain
└── pkg
    └── action
        └── ac001
```

```
awscnfm cl001 ac001 execute
awscnfm cl001 ac001 explain
awscnfm cl001 ac002 ...
```



### Release Versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we would tag
`conformance` with the same version. That way we keep the release and test
lifecycles synchronized.



### Execute Tests

```
awscnfm cl001 ac001 execute
```



### Explain Tests

```
awscnfm cl001 ac001 explain
```



### Tool Version

```
awscnfm version
```
