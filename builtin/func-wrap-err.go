package builtin

import (
	"fmt"
	"runtime"
	"strings"
)

// wraps an error with function name prefix
func FuncWrapErr(err error) error {
	programCounter, _, _, _ := runtime.Caller(1)
	fnName, _ := functionAndPkgName(runtime.FuncForPC(programCounter).Name())
	return fmt.Errorf("%s: %w", fnName, err)
}

// wraps an error with package (base)name and function name prefix. e.g. `jsonfile.Read: ...`
func PackageFuncWrapErr(err error) error {
	programCounter, _, _, _ := runtime.Caller(1)
	fnName, pkgBasename := functionAndPkgName(runtime.FuncForPC(programCounter).Name())
	return fmt.Errorf("%s.%s: %w", pkgBasename, fnName, err)
}

// does really ugly things to extract from Go's stack trace formatted function name the:
// - function name (but without lambda suffixes) and
// - the package base name (= without full package prefix)
func functionAndPkgName(funcFullName string) (string, string) {
	const (
		dotLen   = len(".")
		slashLen = len("/")
	)
	// looks like "github.com/org/project/pkg/mypkg.[(*receiver).]FunctionName[.func1[.1[.2]]]"
	afterSlash := funcFullName
	if slashIndex := strings.LastIndex(funcFullName, "/"); slashIndex != -1 {
		afterSlash = funcFullName[slashIndex+slashLen:]
	}
	firstDot := strings.Index(afterSlash, ".")
	if firstDot == -1 {
		panic("unexpected: no dots at all")
	}
	// "github.com/tiiuae/fog_hyper/pkg/common.TestFoo" => "common"
	pkgBasename := afterSlash[:firstDot]
	// "github.com/tiiuae/fog_hyper/pkg/common.TestFoo" => "TestFoo"
	afterFirstDot := afterSlash[firstDot+dotLen:]

	afterReceiver := 0                                       // assume no receiver
	if hasReceiver := afterFirstDot[0] == '('; hasReceiver { // looks like "(*myReader).Read"?
		afterReceiver = strings.Index(afterFirstDot, ".") + 1
	}

	if dotAfterReceiver := stringsIndexOffset(afterFirstDot, ".", afterReceiver); dotAfterReceiver != -1 { // strip off any lambdas
		return afterFirstDot[:dotAfterReceiver], pkgBasename
	} else {
		return afterFirstDot, pkgBasename
	}
}

// same as `strings.Index()` but only searches after an offset (the offset is inclusive in the returned index result)
func stringsIndexOffset(s string, substr string, offset int) int {
	switch index := strings.Index(s[offset:], substr); index {
	case -1:
		return -1
	default:
		return offset + index
	}
}
