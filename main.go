package main

import (
	"example/src/infra/http"
	"log/slog"
	"os"
	"runtime"
)

func main() {
	slog.Info("Starting example:", "version", runtime.Version(), "PID", os.Getpid())
	http.Run()
}
