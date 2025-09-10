# logger

A simple and lightweight logging library for Go.

## Features

-   Multiple log levels (DEBUG, INFO, WARN, ERROR)
-   Colored log output for better readability
-   Displays timestamp, file name, and line number
-   Easy to set up and use
-   Webhook support (Discord)

## Installation

```bash
go get github.com/bankusy/logger
```

## Usage

Here is a simple example of how to use `logger`:

```go
package main

import (
	"github.com/bankusy/logger/model"
	"github.com/bankusy/logger/model/webhooks"
)

func main() {
	// Set the minimum log level.
	// Only logs with this level or higher will be printed.
	logger.Default("DEBUG")
	logger.Info("This is an info message.")
	logger.Error("This is an error message.")
	logger.Debug("This message will not be printed.")
	logger.Warn("This is a warning message.")
}
```

## Preview

<img width="506" height="58" alt="image" src="https://github.com/user-attachments/assets/2eda5e80-08e5-4754-9ec4-08d203748fea" />

## Log Levels

The following log levels are available:

-   `DEBUG`: For detailed debugging information.
-   `INFO`: For general informational messages.
-   `WARN`: For warnings that might indicate a problem.
-   `ERROR`: For errors that have occurred.

You can set the minimum log level using the `logger.Default()` function. For example, if you set the level to `INFO`, `DEBUG` messages will not be printed.

## Log Format

The log output is formatted as follows:

```
[<LEVEL>][<file>:<line>][<timestamp>] <message>
```

-   `<LEVEL>`: The log level (e.g., INFO, ERROR).
-   `<file>:<line>`: The file name and line number where the log was called.
-   `<timestamp>`: The time when the log was created.
-   `<message>`: The log message.

## Webhooks

`logger` supports sending log messages to webhooks. Currently, Discord is supported.

### Discord

To send log messages to a Discord channel, you need to create a `DiscordNotifier` and add it to the logger.

```go
package main

import (
	"github.com/bankusy/logger/model"
	"github.com/bankusy/logger/model/webhooks"
)

func main() {
	logger.Default("DEBUG")

	discordNotifier := &webhooks.DiscordNotifier{
		WebhookUrl: "YOUR_DISCORD_WEBHOOK_URL",
	}

	logger.AddWebhooks(discordNotifier)

	logger.Info("This message will be sent to Discord.")
}
```
