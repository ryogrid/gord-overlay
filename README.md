# **Gord-Overlay**
- [Gord](https://github.com/taisho6339/gord) is a reference implementation of [Chord protocol](https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf).
- **Gord-Overlay** is a fork of Gord which is enabled to run on overlay network constructed with [gossip-overlay lib](https://github.com/ryogrid/gossip-overlay)
  - Additionaly, implementation of on-memory KVS store functionality is planned

---

## What is Gord-Overlay?
Gord-Overlay is a DHT based distribute key-value store.
Gord-Overlay will start as a gRPC server and your application can access data via gRPC.

## How is it work?
Gord-Overlay is an implementation of [DHT Chord](https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf) and simple key-value store using the DHT.
Chord protocol is an algorithm which extends consistent hashing.
Gord server, using chord protocol, allocates a key to a node in distributed nodes ring.

![chord ring](docs/architecture-1.png) 

In addition, the gord servers communicate with each other via gRPC to synchronize route information.
Then, the server can query via gRPC to resolve the node and communicate with the node.

## Usage
Gord's gRPC server reqires a hostname and port number pair.  
If you specify 127.0.0.1:26000, the server will use 127.0.0.1:26000 to communication between other nodes and use 127.0.0.1:26001 to listen local gRPC request.
Specified hostname and port number pair is internally used as a node identifier, so you need to specify a unique pair for each node.
Additionaly, you need to specify address of a node which is already in the network to join the network except the first node.

```
## Build
make build

## Check how to use this command
./gordctl -h

## Start server
./gordctl -l hostAndPort(required) -n existNodeHostAndPort(optional)
```

## Examples
0. Install [grpcurl](https://github.com/fullstorydev/grpcurl)

1. Start servers
```bash
# start three server processes
docker-compose build && docker-compose up
```

2. Try! 
```bash
# Check successor list for each node
grpcurl -plaintext localhost:26040 server.InternalService/Successors \
&& grpcurl -plaintext localhost:36040 server.InternalService/Successors \
&& grpcurl -plaintext localhost:46040 server.InternalService/Successors

# Check predecessor for each node
grpcurl -plaintext localhost:26040 server.InternalService/Predecessor \
&& grpcurl -plaintext localhost:36040 server.InternalService/Predecessor \
&& grpcurl -plaintext localhost:46040 server.InternalService/Predecessor

# Query
grpcurl -plaintext -d '{"key": "gord1"}' localhost:26041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord1"}' localhost:36041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord1"}' localhost:46041 server.ExternalService/FindHostForKey 

grpcurl -plaintext -d '{"key": "gord2"}' localhost:26041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord2"}' localhost:36041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord2"}' localhost:46041 server.ExternalService/FindHostForKey 

grpcurl -plaintext -d '{"key": "gord"}' localhost:26041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord"}' localhost:36041 server.ExternalService/FindHostForKey \
&& grpcurl -plaintext -d '{"key": "gord"}' localhost:46041 server.ExternalService/FindHostForKey 
```

## How to build
```bash
make build
```

## How to run tests
```bash
make test
```
