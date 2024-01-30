module github.com/ryogrid/gord-overlay

go 1.21.2

//replace github.com/ryogrid/gossip-overlay => ../gossip-overlay

require (
	connectrpc.com/connect v1.14.0
	//github.com/e-dard/netbug v0.0.0-20151029172837-e64d308a0b20
	github.com/ryogrid/gossip-overlay v0.0.13
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	//github.com/e-dard/netbug v0.0.0-20151029172837-e64d308a0b20
	github.com/weaveworks/mesh v0.0.0-20191105120815-58dbcc3e8e63 // indirect
	golang.org/x/net v0.17.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pion/datachannel v1.5.5 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/sctp v1.8.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
