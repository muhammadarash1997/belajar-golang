package main

import (
	"fmt"
	oslog "log"
	"os"

	"github.com/go-kit/log"
)

// GO-KIT LOGGER SUPPORT OUTPUT FORMATS:
// Logfmt)
// JSON

func main() {
	// logging1()
	// logging2()
	// logging3()
	// logging4()
	logging5()
	// logging6()
}

func logging1() {
	fmt.Println("logging1")

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	logger.Log("question", "how old are you?", "answer", 21)
}

func logging2() {
	fmt.Println("logging2")

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "instance_id", 123)

	logger.Log("msg", "starting")

	logger = log.With(logger, "component", "worker")
	logger = log.With(logger, "component", "slacker")

	logger.Log("msg", "running")
}

func logging3() {
	fmt.Println("logging3")

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	oslog.SetOutput(log.NewStdlibAdapter(logger))
	oslog.Print("I sure like pie")
}

func logging4() {
	fmt.Println("logging4")

	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	logger.Log("legacy", true, "msg", "at least it's something")
}

func logging5() {
	fmt.Println("logging5")

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	logger.Log("msg", "hello")
}

func logging6() {
	fmt.Println("logging6")

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	logger.Log("msg", "hello")
}