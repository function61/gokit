// Dynamically embed a build version at compile time, so you can statically access it at runtime.
package dynversion

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

// current user-friendly revision identifier for the currently running software
var Version = "dev"

func IsDevVersion() bool {
	return strings.HasSuffix(Version, "dev")
}

func init() {
	version, err := resolveVersion()
	if err != nil {
		Version = fmt.Sprintf("[Error: %v]", err)
	} else {
		Version = version
	}
}

func resolveVersion() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("unable to read build info")
	}

	setting := func(key string) *debug.BuildSetting { return findSetting(key, info.Settings) } // helper

	rev, timeStr, modified := setting("vcs.revision"), setting("vcs.time"), setting("vcs.modified")
	if rev == nil || timeStr == nil {
		return "", errors.New("no VCS metadata")
	}

	ts, err := time.Parse(time.RFC3339, timeStr.Value)
	if err != nil {
		return "", err
	}

	revisionIdShort := rev.Value[0:8]
	if modified != nil && modified.Value == "true" {
		revisionIdShort += "-dev"
	}

	// https://chronver.org/ -ish (we're using revision ID instead of branch name)
	friendlyRevId := ts.UTC().Format("2006.01.02") + "-" + revisionIdShort
	return friendlyRevId, nil
}

func findSetting(key string, settings []debug.BuildSetting) *debug.BuildSetting {
	for _, setting := range settings {
		if setting.Key == key {
			return &setting
		}
	}

	return nil
}
