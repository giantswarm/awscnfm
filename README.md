[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](https://pkg.go.dev/github.com/giantswarm/awscnfm?tab=overview)



# awscnfm

Conformance test tool for Kubernetes on AWS by Giant Swarm.



## Basic concepts

Tests are organized based on **test plans** containing multiple **actions**.
Each test plan represents a distinct cluster configuration and tested behaviour.
The actions being executed throughout a test plan execution take care of things
like creating the cluster, checking that it functions as expected, and deleting
it.

For example, there could be one test plan for a cluster with one node pool and
HA masters, while another test plan would be for single master and no node
pools. You can either [execute the entire test plan] the entire test plan or a
single action.



## Documentation

Find all [documentation in the `docs`] folder of this repo.



## Release versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we tag `awscnfm` with
the same version. That way we keep the release and test lifecycles synchronized.



[execute the entire test plan]: https://github.com/giantswarm/awscnfm/blob/master/docs/using.md
[documentation in the `docs`]: https://github.com/giantswarm/awscnfm/tree/master/docs
