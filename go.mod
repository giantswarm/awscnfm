module github.com/giantswarm/awscnfm/v12

go 1.15

require (
	github.com/fatih/color v1.12.0
	github.com/giantswarm/apiextensions/v3 v3.26.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/certs/v3 v3.1.1
	github.com/giantswarm/columnize v2.0.2+incompatible
	github.com/giantswarm/k8sclient/v5 v5.11.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/tenantcluster/v3 v3.0.0
	github.com/giantswarm/valuemodifier v0.3.1
	github.com/google/go-cmp v0.5.6
	github.com/jsonmaur/aws-regions/go 7e6567e0cad4
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/spf13/cobra v1.1.3
	golang.org/x/net abc453219eb5
	k8s.io/api v0.18.19
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
	sigs.k8s.io/cluster-api v0.3.17
	sigs.k8s.io/controller-runtime v0.6.4
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
