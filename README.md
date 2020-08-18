[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](https://pkg.go.dev/github.com/giantswarm/awscnfm?tab=overview)

# awscnfm

Conformance test tool for Kubernetes on AWS by Giant Swarm.

## Basic concepts

Tests are organized based on **cluster scopes** containing multiple **actions**. Each cluster scope represents a distinct cluster configuration. The actions take care of things like creating the cluster, checking that it functions as expected, and deleting it.

For example, there could be one cluster scope for a cluster with one node pool and HA masters, while another cluster scope would be for single master and no node pools.

You can either [execute](https://github.com/giantswarm/awscnfm/blob/master/docs/using.md) the entire cluster scope or a single action.

The `awscnfm` major and minor **version** used for testing is supposed to match the major and minor version of the release used by the tenant cluster(s) to test.

## Documentation

Find all documentation in the [`docs`](https://github.com/giantswarm/awscnfm/tree/master/docs) folder of this repo.

## Setup

The first thing you need in order to use `awscnfm` is to have an established
connection to a Control Plane. See the [docs about the environment variables
needed](/docs/environment.md) in order to do that. You also want to get
comfortable with the [project structure](/docs/structure.md) so that you
understand how to execute or explain certain actions for a specific cluster
scope. More about terminology and several basics can also be found in the
[introduction docs](/docs/introduction.md).

## Release versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we tag `awscnfm` with
the same version. That way we keep the release and test lifecycles synchronized.
