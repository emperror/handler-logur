package logur_test

import (
	"fmt"

	"emperror.dev/errors"
	"logur.dev/logur"

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

func newLogurLogger() *errorLogger {
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
