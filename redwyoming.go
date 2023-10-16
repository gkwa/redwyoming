package redwyoming

import (
	"log/slog"
	"os"
)

func Main() {
	slog.Debug("redwyoming", "test", true)

	test1()
}

func test1() {
	tmpDir := os.TempDir()
	slog.Debug("temporary directory", "directory", tmpDir)
}
