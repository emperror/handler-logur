package logur

import (
	"testing"

	"emperror.dev/errors"
	"github.com/goph/logur"
	"github.com/goph/logur/logtesting"
)

type errorsStub struct {
	err     error
	details []interface{}
}

func (e *errorsStub) Error() string          { return e.err.Error() }
func (e *errorsStub) Errors() []error        { return errors.GetErrors(e.err) }
func (e *errorsStub) Details() []interface{} { return e.details }

func TestHandler(t *testing.T) {
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
				Line:  "error",
				Level: logur.Error,
			},
		},
		errors.Combine(
			errors.New("error 1"),
			errors.New("error 2"),
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
	}

	for err, expectedEvents := range tests {
		err, expectedEvents := err, expectedEvents

		t.Run("", func(t *testing.T) {
			logger := logur.NewTestLogger()
			handler := New(logger)

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
}

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
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_test.go:155"},
			},
		},
		errors.Combine(
			errors.New("error 1"),
			errors.New("error 2"),
		): {
			{
				Line:   "error 1",
				Level:  logur.Error,
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_test.go:163"},
			},
			{
				Line:   "error 2",
				Level:  logur.Error,
				Fields: map[string]interface{}{"func": "TestWithStackInfo", "file": "handler_test.go:164"},
			},
		},
	}

	for err, expectedEvents := range tests {
		err, expectedEvents := err, expectedEvents

		t.Run("", func(t *testing.T) {
			logger := logur.NewTestLogger()
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
}
