module github.com/giantswarm/awscnfm/v15

go 1.15

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/fatih/color v1.13.0
	github.com/giantswarm/apiextensions/v3 v3.39.0
	github.com/giantswarm/backoff v1.0.0
	github.com/giantswarm/certs/v3 v3.1.1
	github.com/giantswarm/columnize v2.0.2+incompatible
	github.com/giantswarm/k8sclient/v5 v5.12.0
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v0.6.0
	github.com/giantswarm/tenantcluster/v4 v4.1.0
	github.com/giantswarm/valuemodifier v0.4.0
	github.com/google/go-cmp v0.5.7
	github.com/jsonmaur/aws-regions/v2 v2.3.1
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/spf13/cobra v1.3.0
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d
	golang.org/x/text v0.3.7
	k8s.io/api v0.18.19
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
	sigs.k8s.io/cluster-api v0.4.1
	sigs.k8s.io/controller-runtime v0.6.4
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
