# go-logging

## Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (<https://pkg.go.dev/log>).

[![GoReportCard example](https://goreportcard.com/badge/github.com/senzing/go-logging)](https://goreportcard.com/report/github.com/senzing/go-logging)
[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-logging.svg)](https://pkg.go.dev/github.com/senzing/go-logging)

## Overview

The Senzing go-logging packages use the message number to coordinate aspects of the log message such as
message identification, message text, status, and logging level.

go-logging also allows different formatting options such as JSON or simply terse messages.

go-logging extends the levels of logging to include:
Trace, Debug, Info, Warn, Error, Fatal, and Panic.

go-logging supports "guards",
e.g. IsXxxxx() methods,
to avoid calling a `Log()` method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG `Log()` method when the
logging level is set to INFO.  Guards prevent this.
Example:

```go
 if logger.IsDebug() {
  logger.Debugf("%s", complexProcess())
 }
```

The basic use of senzing/go-logging looks like this:

```go
 import "log"
 import "github.com/senzing/go-logging/messagelogger"

 log.SetFlags(0)
 messageLogger, _ := messagelogger.New()
 messageLogger.Log(1)
```

Output:

```console
 INFO 1:
```

The API documentation and more examples are available at
[pkg.go.dev/github.com/senzing/go-logging](https://pkg.go.dev/github.com/senzing/go-logging)

## Details

The packages of `go-logging` can be though of as belonging to one of the following four groups:

1. **message fields:** `messagedate`, `messagedetails`, `messageduration`, `messageerrors`, `messagid`, `messagelevel`, `messagelocation`, `messagestatus`, `messagetext`, `messagetime`
1. **message format:** `messageformat`
1. **message use:** `messagelogger`
1. **logging:**  `logger`

### Message fields

Packages that manage message fields are:

- `messagedate`
- `messagedetails`
- `messageduration`
- `messageerrors`
- `messagid`
- `messagelevel`
- `messagelocation`
- `messagestatus`
- `messagetext`
- `messagetime`

These packages have a method signature similar to:

```go
 MessageXxxx(messageNumber int, details ...interface{}) (string, error)
```

They receive the message identification number and a series of details.
From this information, they construct the value of the field to be logged.
If the returned string is empty, that field does not appear in the final message.

### Message format

Packages that manage message fields are:

- `messageformat`

These packages have a similar method signature:

```go
 Message(
    date string,
    time string,
    level string,
    location string,
    id string,
    status string,
    text string,
    duration int64,
    errors interface{},
    details interface{}
    ) (string, error)
```

The method receives a value for each field,
which was probably generated by a "Message field" method,
and aggregates them into a string representation.
A message formatter chooses which fields to include and the format of the final message.
The string representation may be JSON, a terse format, or a user-defined format.

### Message use

Packages that use messages are:

- `messagelogger`

"Message use" includes: Logging, Error creation, and simple message generation.
In the case of Logging, a logging level may be set to prevent "low-level" log message from being written to the log.

### Logging

Packages that write log messages are:

- `logger`

This package sit on top of Go's
[log](https://pkg.go.dev/log)
package.
The `logger` package does not replace Go's `log` package.
Rather it puts a
[Decorator pattern](https://en.wikipedia.org/wiki/Decorator_pattern)
on top of Go's `log` package to support concepts such as:

- Log levels of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC
- Guards.  Examples: IsTrace(), IsDebug(), etc.
