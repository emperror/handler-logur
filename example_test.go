package logur_test

import (
	"errors"
	"fmt"

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
