module github.com/giantswarm/awscnfm

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/giantswarm/apiextensions v0.4.17
	github.com/giantswarm/certs/v2 v2.0.0
	github.com/giantswarm/k8sclient/v3 v3.1.2
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/tenantcluster/v2 v2.0.0
	github.com/google/go-cmp v0.5.1
	github.com/spf13/cobra v1.0.0
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
	k8s.io/api v0.17.8
	k8s.io/client-go v0.17.8
	sigs.k8s.io/cluster-api v0.3.7
	sigs.k8s.io/controller-runtime v0.5.8
)
