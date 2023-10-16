package redwyoming

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func Main() {
	slog.Debug("redwyoming", "test", true)

	test1()
	test2()
}

func test1() {
	tmpDir := os.TempDir()
	fmt.Printf("Temporary directory: %s\n", tmpDir)
}

func getTemporaryDirectory() string {
	if runtime.GOOS == "windows" {
		return os.TempDir()
	} else {
		return "/tmp"
	}
}

func test2() {
	tmpDir := getTemporaryDirectory()
	fmt.Printf("Temporary directory: %s\n", tmpDir)
}
