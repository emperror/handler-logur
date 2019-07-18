package logur_test

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/goph/logur"

	logurhandler "emperror.dev/handler/logur"
)

func ExampleNew() {
	logger := logur.NewNoopLogger()
	_ = logurhandler.New(logger)

	// Output:
}

type errorLogger struct{}

func (e *errorLogger) Error(msg string, fields ...map[string]interface{}) {
	fmt.Println(msg)
	if len(fields) > 0 && len(fields[0]) > 0 {
		fmt.Println(fields[0])
	}
}

func newLogurLogger() logur.ErrorLogger {
	return &errorLogger{}
}

func ExampleHandler_Handle() {
	logger := newLogurLogger()
	handler := logurhandler.New(logger)

	err := errors.New("error")

	handler.Handle(err)

	// Output:
	// error
}

func ExampleWithStackInfo() {
	logger := newLogurLogger()
	handler := logurhandler.WithStackInfo(logurhandler.New(logger))

	err := errors.New("error")

	handler.Handle(err)

	// Output:
	// error
	// map[file:example_test.go:48 func:ExampleWithStackInfo]
}
