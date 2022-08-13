package main

import (
	"errors"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	// logging1()
	logging2()
	// logging3()
}

func logging1() {
	logger := log.NewLogfmtLogger(os.Stdout)

	level.Debug(logger).Log("msg", "this message is at the debug level")
	level.Info(logger).Log("msg", "this message is at the info level")
	level.Warn(logger).Log("msg", "this message is at the warn level")
	level.Error(logger).Log("msg", "this message is at the error level")
}

func logging2() {
	// Set up logger with level filter.
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowError())	// minimal is error level will be shown
	logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestampUTC)

	// Use level helpers to log at different levels.
	level.Debug(logger).Log("next item", 17) // filtered
	level.Info(logger).Log("event", "data saved") // filtered
	level.Warn(logger).Log("msg", "this message is at the warn level") // filtered
	level.Error(logger).Log("err", errors.New("bad data"))
}

func logging3() {
	// Set up logger with level filter.
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowInfo())	// minimal is info level will be shown
	logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestampUTC)

	// Use level helpers to log at different levels.
	level.Debug(logger).Log("next item", 17) // filtered
	level.Info(logger).Log("event", "data saved")
	level.Warn(logger).Log("msg", "this message is at the warn level")
	level.Error(logger).Log("err", errors.New("bad data"))
}
