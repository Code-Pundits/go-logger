package main

import (
	logging "github.com/Code-Pundits/go-logger"
)

func main() {
	logger := logging.
		NewLogger().
		WithLevel(logging.InfoLevel).
		WithTransports(
			logging.NewStdOutTransport(logging.StdOutTransportConfig{Level: logging.InfoLevel}),
		).
		WithDefaults(
			&logging.FieldPair{Name: "Component", Value: "hello-world-service"},
		)

	logger.Info("Hello World!")
}
