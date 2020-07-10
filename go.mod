module github.com/giantswarm/awscnfm

go 1.14

require (
	github.com/giantswarm/apiextensions v0.4.8
	github.com/giantswarm/certs/v2 v2.0.0
	github.com/giantswarm/k8sclient/v3 v3.1.1
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/tenantcluster/v2 v2.0.0
	github.com/spf13/cobra v1.0.0
	k8s.io/api v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/cluster-api v0.3.6
	sigs.k8s.io/controller-runtime v0.5.2
)
