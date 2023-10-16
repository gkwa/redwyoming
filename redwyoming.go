package redwyoming

import (
	"log/slog"
	"os"
)

func Main() {
	tmpDir := os.TempDir()
	slog.Debug("temporary directory", "directory", tmpDir)
}
