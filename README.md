[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](https://pkg.go.dev/github.com/giantswarm/awscnfm?tab=overview)

# awscnfm

Conformance test utility for AWS.



### Docs

* [Completion](/docs/completion.md)
* [Structure](/docs/structure.md)



### Setup

As of now you need to start your own Tenant Cluster so that you can execute the
implemented conformance tests against it. Cluster creation can be another action
supported in the test suite at a later point in time.



### Release Versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we tag `awscnfm` with
the same version. That way we keep the release and test lifecycles synchronized.
