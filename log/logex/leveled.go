package logex

import (
	"fmt"
	"log"
)

// purposefully minimal # of log levels. I kind of agree with this post ..:
//
//	https://dave.cheney.net/2015/11/05/lets-talk-about-logging
//
// .. with the exception, that Error levels are justified
type Leveled struct {
	Debug *log.Logger // usually hidden, maybe noisy, but can be enabled for debugging deep problems
	Info  *log.Logger // something interesting happened
	Warn  *log.Logger // something unexpected happened, but not yet jeopardizing the application
	Error *log.Logger // something exceptional operations usually should look at: like I/O (network/disk) failed, service call failed, security problem etc.

	// convenience for if some component needs raw (= non-leveled) Logger
	Original *log.Logger
}

func Levels(parent *log.Logger) *Leveled {
	return &Leveled{
		Debug: Prefix("[DEBUG]", parent),
		Info:  Prefix("[INFO]", parent),
		Warn:  Prefix("[WARN]", parent),
		Error: Prefix("[ERROR]", parent),

		Original: parent,
	}
}

// sometimes you may need to pipe log output from logging interfaces with more levels than
// what we have declared here. prefer uppercase: Prefix(CustomLevelPrefix("WARN"))
func CustomLevelPrefix(levelName string) string {
	return fmt.Sprintf("[%s]", levelName)
}
