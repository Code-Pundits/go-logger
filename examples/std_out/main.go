package main

import log "github.com/Code-Pundits/go-logger"

func main() {
	logger := log.
		NewLogger().
		WithLevel(log.InfoLevel).
		WithTransports(
			log.NewStdOutTransport(log.StdOutTransportConfig{Level: log.InfoLevel}),
		).
		WithDefaults(
			&log.FieldPair{Name: "Component", Value: "hello-world-service"},
		)

	logger.Info("here we are!")
}
