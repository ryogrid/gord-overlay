module github.com/ryogird/gord-overlay

go 1.18

replace github.com/ryogrid/gord-overlay => ./

require (
	connectrpc.com/connect v1.14.0
	github.com/golang/protobuf v1.5.3
	github.com/ryogrid/gord-overlay v0.0.0-20240101010810-e130a228ccb9
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.2.2
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.32.0
)
