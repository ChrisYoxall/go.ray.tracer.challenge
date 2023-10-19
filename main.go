package main

import (
	"log/slog"
)

func main() {

	// Should I play with slog?
	slog.Info("Hello World!", slog.Int("value", 123), slog.String("ip", "127.0.0.1"))

}
