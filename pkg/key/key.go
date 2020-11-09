package key

import (
	"fmt"
	"strings"

	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

const (
	// Credential is the default credential we use for most of our conformance
	// test clusters. These credentials define which AWS Account to use.
	Credential = "credential-default"
	// Organization is the Giant Swarm specific organization we create our
	// conformance test clusters in.
	Organization = "giantswarm"
)

const (
	KiamTestNamespace = "kube-system"
)

func APIEndpoint(id string, base string) string {
	return fmt.Sprintf("api.%s.k8s.%s", id, base)
}

func DomainFromHost(h string) string {
	h = strings.Replace(h, "https://", "", 1)
	h = strings.Replace(h, "g8s.", "", 1)
	h = strings.Replace(h, ":443", "", 1)
	return h
}

func KiamTestJobName() string {
	return fmt.Sprintf("%s-kiam-test", project.Name())
}

func KiamTestNetPolName() string {
	return fmt.Sprintf("%s-kiam-test", project.Name())
}

func RegionFromHost(h string) string {
	h = strings.Split(DomainFromHost(h), ".")[1]
	// Domain in giraffe is giraffe.pek.aws.k8s.adidas.com.cn which does not follow the convention
	if h == "pek" {
		h = "cn-north-1"
	}
	return h
}
