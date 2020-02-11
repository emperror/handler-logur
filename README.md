# Logur error handler

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/handler-logur/CI?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/handler-logur?style=flat-square)](https://codecov.io/gh/emperror/handler-logur)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/logur?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/logur)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-logur.svg)](https://golangci.com/r/github.com/emperror/handler-logur)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)](https://github.com/emperror/errors)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/emperror.dev/handler/logur)

**Error handler using [Logur](https://github.com/goph/logur).**


## Installation

```bash
go get emperror.dev/handler/logur
```


## Usage

```go
package main

import (
	"github.com/goph/logur/adapters/logrusadapter"
	"github.com/sirupsen/logrus"

	logurhandler "emperror.dev/handler/logur"
)

func main() {
	logger := logrusadapter.New(logrus.New())
	handler := logurhandler.New(logger)
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
