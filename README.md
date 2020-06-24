[![CircleCI](https://circleci.com/gh/giantswarm/awscnfm.svg?style=shield)](https://circleci.com/gh/giantswarm/awscnfm)
[![GoDoc](https://godoc.org/github.com/giantswarm/awscnfm?status.svg)](http://godoc.org/github.com/giantswarm/awscnfm)

# awscnfm



### Structure

The tool should be structured so that it can execute any action against our
Tenant Clusters. There might be many cluster definitions for which have many
actions each. The command line tool layout could semantically look like below.

```
awscnfm cluster 001 action 001 execute
awscnfm cluster 001 action 001 explain
awscnfm cluster 001 action 001 ...
```



### Release Versions

The code used for testing conformity is aligned with every Giant Swarm release
version. When publishing a `v11.4.0` Giant Swarm release we would tag
`conformance` with the same version. That way we keep the release and test
lifecycles synchronized.



### Execute Tests

```
awscnfm cluster 001 action 001 execute
```



### Explain Tests

```
awscnfm cluster 001 action 001 explain
```



### Tool Version

```
awscnfm version
```
