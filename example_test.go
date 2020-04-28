package logur_test

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"logur.dev/logur"

	logurhandler "emperror.dev/handler/logur"
)

func ExampleNew() {
	logger := &logur.NoopLogger{}
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

func (e *errorLogger) ErrorContext(ctx context.Context, msg string, fields ...map[string]interface{}) {
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

func ExampleHandler_HandleContext() {
	logger := newLogurLogger()
	handler := logurhandler.New(logger)

	ctx := context.Background()
	err := errors.New("error")

	handler.HandleContext(ctx, err)

	// Output:
	// error
}
