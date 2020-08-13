# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [12.1.1] - 2020-08-12

## [12.1.0] - 2020-08-12

### Added

* Add code generation.
* Add first documentation.
* Add first cluster scope and actions.
* New action to verifiy if host network pods matches with k8scloudconfig on master nodes
* New action to verifiy if host network pods matches with k8scloudconfig on worker nodes

### Fixed

* Include namespace to delete cluster CR's
* Drop tenant cluster initialization in cleaner verification action

[Unreleased]: https://github.com/giantswarm/awscnfm/compare/v12.1.1...HEAD
[12.1.1]: https://github.com/giantswarm/awscnfm/compare/v12.1.0...v12.1.1
[12.1.0]: https://github.com/giantswarm/awscnfm/releases/tag/v12.1.0
