package dynversion

// replaced dynamically at build time from function61/buildkit-golang/build-common.sh
var Version = "dev"

func IsDevVersion() bool {
	return Version == "dev"
}
