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
	github.com/giantswarm/micrologger v1.0.0
	github.com/giantswarm/tenantcluster/v4 v4.1.0
	github.com/giantswarm/valuemodifier v0.4.0
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/google/go-cmp v0.5.7
	github.com/google/uuid v1.3.0 // indirect
	github.com/jsonmaur/aws-regions/v2 v2.3.1
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/spf13/cobra v1.3.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	golang.org/x/sys v0.0.0-20220608164250-635b8c9b7f68 // indirect
	golang.org/x/text v0.3.7
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
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
