package release

type ComponentsContainer interface {
	Latest() map[string]string
	Previous() map[string]string
}

type Resolver interface {
	Components() ComponentsContainer
	Version() VersionContainer
}

// VersionContainer implements Latest and Previous which return their respective
// semver level version depending on the MAJOR, MINOR and PATCH level
// implementations. Consider a project version of the awscnfm tool of v12.2.x
// and the following releases are available on the Control Plane.
//
//     MAJOR               MINOR               PATCH
//
//     ...
//     v11.3.0
//     v11.5.0 previous
//     v11.5.3 latest      ...
//     ...                 v12.0.1
//                         v12.0.2 previous
//                         v12.1.0 latest      ...
//                         ...                 v12.2.2
//                                             v12.2.4 previous
//                                             v12.2.5 latest
//                                             ...
//
// In the described scenario above example behaviour should be as follows.
//
//     * the MAJOR level implementation of VersionContainer.Latest returns v11.5.3
//     * the MINOR level implementation of VersionContainer.Latest returns v12.1.0
//     * the PATCH level implementation of VersionContainer.Latest returns v12.2.5
//
type VersionContainer interface {
	Latest() string
	Previous() string
}
