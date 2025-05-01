# Golang auto slog (structured logger)

[![Go Reference](https://pkg.go.dev/badge/github.com/iguanesolutions/auto-slog.svg)](https://pkg.go.dev/github.com/iguanesolutions/auto-slog)

This lib allows to simply create a [Golang structured logger](https://go.dev/blog/slog) with its handler automatically selected:

* If the output is a terminal, a standard [text handler](https://pkg.go.dev/log/slog#TextHandler) will be used
* If the program has been started by systemd, a custom [journald handler](https://pkg.go.dev/github.com/iguanesolutions/go-systemd/v5@v5.2.0/journald/slog) will be used
* Otherwise a standard [JSON handler](https://pkg.go.dev/log/slog#JSONHandler) will be used

## Installation

```bash
go get -u github.com/iguanesolutions/auto-slog
```

## Usage

```go
package main

import (
	"log/slog"

	autoslog "github.com/iguanesolutions/auto-slog"
)

func main() {
	logger := autoslog.NewLogger(autoslog.LogLevel("INFO"))
	logger.Info("Hello !", slog.Any("key", "value"))
}
```
