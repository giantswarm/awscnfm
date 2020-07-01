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

func APIEndpoint(id string, base string) string {
	return fmt.Sprintf("api.%s.k8s.%s", id, base)
}

func GeneratedWithPrefix(s string) string {
	return fmt.Sprintf("%s%s", GeneratePrefix, s)
}

func HasGeneratedPrefix(s string) bool {
	return strings.HasPrefix(s, GeneratePrefix)
}
