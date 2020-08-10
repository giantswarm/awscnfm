[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](https://pkg.go.dev/github.com/giantswarm/awscnfm?tab=overview)



# awscnfm

Conformance test tool for AWS.



### Docs

* [Completion](/docs/completion.md)
* [Environment](/docs/environment.md)
* [Generation](/docs/generation.md)
* [Introduction](/docs/introduction.md)
* [Structure](/docs/structure.md)



### Setup

The first thing you need in order to use `awscnfm` is to have an established
connection to a Control Plane. See the [docs about the environment variables
needed](/docs/environment.md) in order to do that. You also want to get
comfortable with the [project structure](/docs/structure.md) so that you
understand how to execute or explain certain actions for a specific cluster
scope. More about terminology and several basics can also be found in the
[introduction docs](/docs/introduction.md).



### Release Versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we tag `awscnfm` with
the same version. That way we keep the release and test lifecycles synchronized.
