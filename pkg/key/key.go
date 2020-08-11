package key

import (
	"fmt"
	"strings"
)

const (
	// GeneratePrefix is a file name prefix we add to generated files by
	// convention. This should show everyone that this file should not be
	// modified as any changes will be overwritten on next generation anyway.
	GeneratePrefix = "zz_generated."
)

const (
	// Credential is the default credential we use for most of our conformance
	// test clusters. These credentials define which AWS Account to use.
	Credential = "credential-default"
	// Organization is the Giant Swarm specific organization we create our
	// conformance test clusters in.
	Organization = "giantswarm"
	// Release is the latest release version we use as base for conformance
	// testing. This version changes with each new release so that we stay
	// synchronized with the latest AWS release we publish.
	Release = "12.1.0"
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

func GeneratedWithPrefix(s string) string {
	return fmt.Sprintf("%s%s", GeneratePrefix, s)
}

func HasGeneratedPrefix(s string) bool {
	return strings.Contains(s, GeneratePrefix)
}

func RegionFromHost(h string) string {
	h = strings.Split(DomainFromHost(h), ".")[1]
	return h
}
