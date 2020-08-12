package project

var (
	description = "Conformance test utility for AWS."
	gitSHA      = "n/a"
	name        = "awscnfm"
	source      = "https://github.com/giantswarm/awscnfm"
	version     = "0.0.1"
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
