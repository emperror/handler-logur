# Logur error handler

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/handler-logur/CI?style=flat-square)](https://github.com/emperror/handler-logur/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/handler-logur?style=flat-square)](https://codecov.io/gh/emperror/handler-logur)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/logur?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/logur)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/emperror.dev/handler/logur)


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
