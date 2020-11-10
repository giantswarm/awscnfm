package project

var (
	description = "Conformance test utility for AWS."
	gitSHA      = "n/a"
	name        = "awscnfm"
	source      = "https://github.com/giantswarm/awscnfm"
	// version is synchronized with the latest Giant Swarm release on the major
	// and minor level so that we know which version we use as base for
	// conformance testing. Please make sure that at least major and minor
	// levels of this version aligns with each new release so that we stay
	// synchronized with the latest AWS release we publish.
	version = "12.6.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
