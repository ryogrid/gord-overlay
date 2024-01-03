module github.com/ryogird/gord-overlay

go 1.14

replace github.com/ryogrid/gord-overlay => ./

require (
	github.com/golang/protobuf v1.5.3
	github.com/ryogrid/gord-overlay v0.0.0-20240101010810-e130a228ccb9 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.2.2
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.32.0
)
