package project

var (
	description = "Conformance test utility for AWS."
	gitSHA      = "n/a"
	name        = "awscnfm"
	source      = "https://github.com/giantswarm/awscnfm"
	// version is synchronnized with the latest Giant Swarm release version so
	// that we know which version we use as base for conformance testing. Please
	// make sure that this version changes with each new release so that we stay
	// synchronized with the latest AWS release we publish.
	version = "12.1.1"
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
