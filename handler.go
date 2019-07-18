// Package logur provides an error handler using a Logur compatible logger.
package logur

import (
	"fmt"

	"emperror.dev/errors"
	"emperror.dev/errors/utils/keyval"
	"github.com/goph/logur"
)

// Handler logs errors using a Logur compatible logger.
type Handler struct {
	logger ErrorLogger

	enableStackInfo bool
}

// ErrorLogger is a subset of the Logur logger interface used for error logging.
type ErrorLogger interface {
	// Error logs an Error event.
	//
	// Critical events that require immediate attention.
	Error(msg string, fields ...map[string]interface{})
}

// New returns a new Handler.
func New(logger ErrorLogger) *Handler {
	if logger == nil {
		logger = logur.NewNoopLogger()
	}

	return &Handler{
		logger: logger,
	}
}

// WithStackInfo enables annotating every error passing through the handler with the
// function name and file line of the stack trace's top frame (if one is found).
func WithStackInfo(handler *Handler) *Handler {
	handler.enableStackInfo = true

	return handler
}

// Handle records an error event and forwards it to the underlying logger.
func (h *Handler) Handle(err error) {
	if err == nil {
		return
	}

	fields := make(map[string]interface{})

	// Extract details from the error
	if details := errors.GetDetails(err); len(details) > 0 {
		fields = keyval.ToMap(details)
	}

	if errs := getErrors(err); len(errs) > 1 || errs[0] == err {
		for _, e := range errs {
			// Extract details from the error
			details := errors.GetDetails(e)
			f := make(map[string]interface{}, len(fields)+len(details)/2)
			for key, value := range fields {
				f[key] = value
			}

			if len(details) > 0 {
				fields := keyval.ToMap(details)

				for key, value := range fields {
					f[key] = value
				}
			}

			h.logError(e, f)
		}
	} else {
		h.logError(err, fields)
	}
}

// fields is always copied when multiple errors are detected,
// so we are free to modify it
func (h *Handler) logError(err error, fields map[string]interface{}) {
	if h.enableStackInfo {
		var stackTracer interface{ StackTrace() errors.StackTrace }
		if errors.As(err, &stackTracer) {
			stackTrace := stackTracer.StackTrace()

			if len(stackTrace) > 0 {
				frame := stackTrace[0]

				fields["func"] = fmt.Sprintf("%n", frame)
				fields["file"] = fmt.Sprintf("%v", frame)
			}
		}
	}

	h.logger.Error(err.Error(), fields)
}

func getErrors(err error) []error {
	if eg, ok := err.(interface{ Errors() []error }); ok {
		errors := eg.Errors()
		result := make([]error, len(errors))
		copy(result, errors)
		return result
	}

	return errors.GetErrors(err)
}
