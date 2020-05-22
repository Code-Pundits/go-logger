# Go Logger

Golang Logger

## Usage

### Start using it
Download go-logger for Go by using:
```sh
$ go get github.com/Code-Pundits/go-logger
```
Import following in your code:
```go
import "github.com/Code-Pundits/go-logger"
```

### Canonical example:

```go
package main

import (
	"github.com/Code-Pundits/go-logger"
)

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

	logger.Info("Hello world")
	// {"timestamp":"2020-05-22T15:06:22-05:00","component":"api-gateway"}
}
```