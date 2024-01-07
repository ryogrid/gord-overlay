module github.com/ryogird/gord-overlay

go 1.14

replace github.com/ryogrid/gord-overlay => ./

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/ryogrid/gord-overlay v0.0.0-20240101010810-e130a228ccb9 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)
