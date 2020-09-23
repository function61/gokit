// Same as dynversion, but for pre-compilation stage (e.g. code generation stage) accesses
// the version via ENV var.
package precompilationversion

import (
	"os"
)

func PreCompilationVersion() string {
	friendlyRevId := os.Getenv("FRIENDLY_REV_ID")
	if friendlyRevId == "" {
		friendlyRevId = "dev"
	}

	return friendlyRevId
}
