module github.com/ryogird/gord-overlay

go 1.18

replace github.com/ryogrid/gord-overlay => ./

require (
	github.com/go-faster/errors v0.7.1
	github.com/golang/protobuf v1.5.3
	github.com/ogen-go/ogen v0.81.1
	github.com/ryogrid/gord-overlay v0.0.0-20240101010810-e130a228ccb9
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)
