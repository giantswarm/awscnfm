module github.com/giantswarm/awscnfm/v12

go 1.15

require (
	github.com/fatih/color v1.10.0
	github.com/giantswarm/apiextensions/v3 v3.14.1
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/certs/v3 v3.1.0
	github.com/giantswarm/columnize v2.0.2+incompatible
	github.com/giantswarm/k8sclient/v5 v5.0.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/tenantcluster/v3 v3.0.0
	github.com/giantswarm/valuemodifier v0.3.0
	github.com/google/go-cmp v0.5.4
	github.com/jsonmaur/aws-regions/go v0.0.0-20200521181458-43baf1be9a5a
	github.com/spf13/cobra v1.1.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	sigs.k8s.io/cluster-api v0.3.11
	sigs.k8s.io/controller-runtime v0.8.0
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
