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

### Std Out Example:

```go
package main

import (
	"github.com/Code-Pundits/go-logger"
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

	logger.Info("hello world!") 
	// {"timestamp":"2020-05-23T19:26:59-05:00","severity":"info","component":"hello-world-service","message":"hello world!"}
}
```