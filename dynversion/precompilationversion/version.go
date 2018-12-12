package precompilationversion

import (
	"os"
)

// we might need access to the version in situations where project has not been compiled yet,
// for example during code generation
func PreCompilationVersion() string {
	friendlyRevId := os.Getenv("FRIENDLY_REV_ID")
	if friendlyRevId == "" {
		friendlyRevId = "dev"
	}

	return friendlyRevId
}
