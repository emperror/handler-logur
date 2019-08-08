// +build go1.12

package logur_test

import (
	"emperror.dev/errors"

	logurhandler "emperror.dev/handler/logur"
)

// Maps are printed in key-sorted order as of Go 1.12
// See https://golang.org/doc/go1.12#fmt
func ExampleWithStackInfo() {
	logger := newLogurLogger()
	handler := logurhandler.WithStackInfo(logurhandler.New(logger))

	err := errors.New("error")

	handler.Handle(err)

	// Output:
	// error
	// map[file:example_go1_12_test.go:17 func:ExampleWithStackInfo]
}
