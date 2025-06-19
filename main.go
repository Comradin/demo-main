package main

import (
	"log/slog"
	"os"

	"demo-main/pkg/foo"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("We are live")

	foobar := foo.New("baz")
	logger.Info("Created FooBar", "content", foobar.Content)
}
