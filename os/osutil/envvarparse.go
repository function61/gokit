package osutil

import (
	"regexp"
)

// this package exists because, for example Docker is so lame that its APIs returns
// ENV vars in serialized form ("key=value")

var envParseRe = regexp.MustCompile("^([^=]+)=(.*)")

func ParseEnvs(serialized string) (string, string) {
	matches := envParseRe.FindStringSubmatch(serialized)
	if matches == nil {
		return "", ""
	}

	return matches[1], matches[2]
}
