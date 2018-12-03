package logex

import (
	"log"
)

type Leveled struct {
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

func Levels(parent *log.Logger) *Leveled {
	return &Leveled{
		Debug: Prefix("[DEBUG]", parent),
		Info:  Prefix("[INFO]", parent),
		Error: Prefix("[ERROR]", parent),
	}
}
