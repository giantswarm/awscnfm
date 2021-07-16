# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Changed

- Shift organization to `conformance-testing`.

## [15.0.3] - 2021-07-16

### Added

- Aliyun registry.

## [15.0.2] - 2021-06-17

### Changed

- Updated expected worker pod list to include aws-ebs-csi-driver.

## [15.0.1] - 2021-06-17

### Changed

- Updated expected app list to include aws-ebs-csi-driver.

## [15.0.0] - 2021-06-04

### Added

- Add action to create EBS Volume job to test CSI driver.
- Add action to verify EBS Volume job to test CSI driver.
- Add action to delete EBS Volume job to test CSI driver.

## [14.2.1] - 2021-05-28

## [14.1.1] - 2021-04-13

## [13.0.1] - 2021-01-08

## [13.0.0] - 2020-12-11

## [12.7.0] - 2020-11-19

### Added

- Add support for `AWSCNFM_UPDATE_RELEASEVERSION`.
- Add action to create network policy test resources.
- Add action to create curl jobs to test network policy.
- Add action to verify curl jobs to test network policy.
- Add action to delete curl jobs to test network policy.
- Add action to delete network policy test resources.
- Add actions to test network policy to plan 001.

### Changed

- Renamed `AWSCNFM_RELEASEVERSION` to `AWSCNFM_CREATE_RELEASEVERSION`.


## [12.6.0] - 2020-11-10



### Changed

- Refactored command line structure to make actions reusable.



## [12.5.0] - 2020-10-19

### Fixed

- Fix giraffe domain detection

### Added

- Added `cl004` for testing NetworkPools.
- Inserted `ac008` for testing kiam app to `cl001`.
- Added Dockerfile.

## [12.1.4] - 2020-09-16

### Added

- Add cluster scope `cl002` for patch upgrade tests.

### Changed

- Tuning backoff time for planner actions.
- Modified the documentation and main README.
- Changed the default description for cluster scopes slightly, for consistent
  use of the "cluster scope" term.
- Updated backward incompatible Kubernetes dependencies to v1.18.5.
- Enable info logging for executing test plans.

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

[Unreleased]: https://github.com/giantswarm/awscnfm/compare/v15.0.3...HEAD
[15.0.3]: https://github.com/giantswarm/awscnfm/compare/v15.0.2...v15.0.3
[15.0.2]: https://github.com/giantswarm/awscnfm/compare/v15.0.1...v15.0.2
[15.0.1]: https://github.com/giantswarm/awscnfm/compare/v15.0.0...v15.0.1
[15.0.0]: https://github.com/giantswarm/awscnfm/compare/v14.2.1...v15.0.0
[14.2.1]: https://github.com/giantswarm/awscnfm/compare/v14.1.1...v14.2.1
[14.1.1]: https://github.com/giantswarm/awscnfm/compare/v13.0.1...v14.1.1
[13.0.1]: https://github.com/giantswarm/awscnfm/compare/v13.0.0...v13.0.1
[13.0.0]: https://github.com/giantswarm/awscnfm/compare/v12.7.0...v13.0.0
[12.7.0]: https://github.com/giantswarm/awscnfm/compare/v12.6.0...v12.7.0
[12.6.0]: https://github.com/giantswarm/awscnfm/compare/v12.5.0...v12.6.0
[12.5.0]: https://github.com/giantswarm/awscnfm/compare/v12.1.4...v12.5.0
[12.1.4]: https://github.com/giantswarm/awscnfm/compare/v12.1.1...v12.1.4
[12.1.1]: https://github.com/giantswarm/awscnfm/compare/v12.1.0...v12.1.1
[12.1.0]: https://github.com/giantswarm/awscnfm/releases/tag/v12.1.0
