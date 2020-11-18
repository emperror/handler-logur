package logur

import (
	"context"
	"testing"

	"emperror.dev/errors"
	"logur.dev/logur"
	"logur.dev/logur/logtesting"
)

func TestWithStackInfo(t *testing.T) {
	tests := map[error][]logur.LogEvent{
		errors.NewPlain("error"): {
			{
				Line:  "error",
				Level: logur.Error,
			},
		},
		errors.Combine(
			errors.NewPlain("error 1"),
			errors.NewPlain("error 2"),
		): {
			{
				Line:  "error 1",
				Level: logur.Error,
			},
			{
				Line:  "error 2",
				Level: logur.Error,
			},
		},
		errors.WithDetails(errors.NewPlain("error"), "key", "value"): {
			{
				Line:   "error",
				Level:  logur.Error,
				Fields: map[string]interface{}{"key": "value"},
			},
		},
		&errorsStub{
			err: errors.Combine(
				errors.WithDetails(errors.NewPlain("error 1"), "key", "value", "key 2", "value 2"),
				errors.NewPlain("error 2"),
			),
			details: []interface{}{"key", "parent value"},
		}: {
			{
				Line:   "error 1",
				Level:  logur.Error,
				Fields: map[string]interface{}{"key": "value", "key 2": "value 2"},
			},
			{
				Line:   "error 2",
				Level:  logur.Error,
				Fields: map[string]interface{}{"key": "parent value"},
			},
		},
		errors.New("error"): {
			{
				Line:   "error",
				Level:  logur.Error,
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_stack_test.go:58"},
			},
		},
		errors.Combine(
			errors.New("error 1"),
			errors.New("error 2"),
		): {
			{
				Line:   "error 1",
				Level:  logur.Error,
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_stack_test.go:66"},
			},
			{
				Line:   "error 2",
				Level:  logur.Error,
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_stack_test.go:67"},
			},
		},
	}

	t.Run("Handle", func(t *testing.T) {
		for err, expectedEvents := range tests {
			err, expectedEvents := err, expectedEvents

			t.Run("", func(t *testing.T) {
				logger := &logur.TestLoggerFacade{}
				handler := WithStackInfo(New(logger))

				handler.Handle(err)

				if got, want := logger.Count(), len(expectedEvents); got != want {
					t.Fatalf("recorded %d events, but expected %d", got, want)
				}

				events := logger.Events()

				for i, expectedEvent := range expectedEvents {
					logtesting.AssertLogEventsEqual(t, expectedEvent, events[i])
				}
			})
		}
	})

	t.Run("HandleContext", func(t *testing.T) {
		for err, expectedEvents := range tests {
			err, expectedEvents := err, expectedEvents

			t.Run("", func(t *testing.T) {
				logger := &logur.TestLoggerFacade{}
				handler := WithStackInfo(New(logger))

				handler.HandleContext(context.Background(), err)

				if got, want := logger.Count(), len(expectedEvents); got != want {
					t.Fatalf("recorded %d events, but expected %d", got, want)
				}

				events := logger.Events()

				for i, expectedEvent := range expectedEvents {
					logtesting.AssertLogEventsEqual(t, expectedEvent, events[i])
				}
			})
		}
	})
}
