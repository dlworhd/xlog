# xlog

A simple and lightweight logging library for Go.

## Features

- Multiple log levels (DEBUG, INFO, WARN, ERROR)
- Colored log output for better readability
- Displays timestamp, file name, and line number
- Easy to set up and use

## Installation

```bash
go get github.com/dlworhd/xlog
```

## Usage

Here is a simple example of how to use `xlog`:

```go
package main

import (
	"github.com/dlworhd/xlog/xlog"
)

func main() {
	// Set the minimum log level.
	// Only logs with this level or higher will be printed.
	xlog.Default("DEBUG")
	xlog.Info("This is an info message.")
	xlog.Error("This is an error message.")
	xlog.Debug("This message will not be printed.")
	xlog.Warn("This is a warning message.")
}
```
## Preview
<img width="393" height="55" alt="image" src="https://github.com/user-attachments/assets/f214306f-b7f6-41af-8e35-f9f946751bd5" />

## Log Levels

The following log levels are available:

- `DEBUG`: For detailed debugging information.
- `INFO`: For general informational messages.
- `WARN`: For warnings that might indicate a problem.
- `ERROR`: For errors that have occurred.

You can set the minimum log level using the `xlog.Default()` function. For example, if you set the level to `INFO`, `DEBUG` messages will not be printed.

## Log Format

The log output is formatted as follows:

```
[<LEVEL>][<file>:<line>][<timestamp>] <message>
```

- `<LEVEL>`: The log level (e.g., INFO, ERROR).
- `<file>:<line>`: The file name and line number where the log was called.
- `<timestamp>`: The time when the log was created.
- `<message>`: The log message.
