module github.com/giantswarm/awscnfm/v15

go 1.15

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/fatih/color v1.13.0
	github.com/giantswarm/apiextensions-application v0.3.1
	github.com/giantswarm/apiextensions/v6 v6.0.0
	github.com/giantswarm/backoff v1.0.0
	github.com/giantswarm/certs/v4 v4.0.0
	github.com/giantswarm/columnize v2.0.2+incompatible
	github.com/giantswarm/k8sclient/v7 v7.0.1
	github.com/giantswarm/k8smetadata v0.10.1
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v0.6.0
	github.com/giantswarm/release-operator/v3 v3.2.0
	github.com/giantswarm/tenantcluster/v6 v6.0.0
	github.com/giantswarm/valuemodifier v0.4.0
	github.com/google/go-cmp v0.5.7
	github.com/jsonmaur/aws-regions/v2 v2.3.1
	github.com/spf13/cobra v1.3.0
	golang.org/x/net v0.17.0
	golang.org/x/text v0.13.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/cluster-api v1.0.4
	sigs.k8s.io/controller-runtime v0.10.3
)

replace (
	github.com/Microsoft/hcsshim v0.8.7 => github.com/Microsoft/hcsshim v0.8.10
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v1.0.4
)
