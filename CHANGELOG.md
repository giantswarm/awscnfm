# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Added

- Added Dockerfile.

### Changed

- Tuning backoff time for planner actions.
- Modified the documentation and main README.
- Changed the default description for cluster scopes slightly, for consistent
  use of the "cluster scope" term.
- Updated backward incompatible Kubernetes dependencies to v1.18.5.

### Fixed

- Make release management work for test releases.

## [12.1.1] - 2020-08-12

## [12.1.0] - 2020-08-12

### Added

* Add code generation.
* Add first documentation.
* Add first cluster scope and actions.
* New action to verifiy if host network pods matches with k8scloudconfig on master nodes.
* New action to verifiy if host network pods matches with k8scloudconfig on worker nodes.
* Add pod names we expect when running host network tests for master and worker node.

### Fixed

* Include namespace to delete cluster CR's.
* Skip aws-cni-restarter for host network test on master.

[Unreleased]: https://github.com/giantswarm/awscnfm/compare/v12.1.1...HEAD
[12.1.1]: https://github.com/giantswarm/awscnfm/compare/v12.1.0...v12.1.1
[12.1.0]: https://github.com/giantswarm/awscnfm/releases/tag/v12.1.0
