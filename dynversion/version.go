// Dynamically embed a build version at compile time, so you can statically access it at runtime.
package dynversion

// !!!!!! WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING !!!!!!
//
// this location is deprecated - please migrate to one in app/dynversion
//
// !!!!!! WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING WARNING !!!!!!

// replaced dynamically at build time from function61/buildkit-golang/build-common.sh
var Version = "dev"

func IsDevVersion() bool {
	return Version == "dev"
}
