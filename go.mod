module github.com/giantswarm/awscnfm

go 1.14

require (
	github.com/fatih/color v1.9.0
	github.com/ghodss/yaml v1.0.0
	github.com/giantswarm/apiextensions v0.4.20
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/certs/v2 v2.0.0
	github.com/giantswarm/columnize v2.0.2+incompatible
	github.com/giantswarm/k8sclient/v3 v3.1.2
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/tenantcluster/v2 v2.0.0
	github.com/google/go-cmp v0.5.1
	github.com/jsonmaur/aws-regions/go v0.0.0-20200521181458-43baf1be9a5a
	github.com/spf13/cobra v1.0.0
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	k8s.io/api v0.17.8
	k8s.io/apimachinery v0.17.8
	k8s.io/client-go v0.17.8
	sigs.k8s.io/cluster-api v0.3.8
	sigs.k8s.io/controller-runtime v0.5.9
)
