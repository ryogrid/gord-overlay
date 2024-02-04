# **Gord-Overlay**
- [Gord](https://github.com/taisho6339/gord) is a reference implementation of [Chord protocol](https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf).
- **Gord-Overlay** is a fork of Gord which is enabled to run on overlay network constructed with [gossip-overlay lib](https://github.com/ryogrid/gossip-overlay)
  - Additionaly, offers on-memory KVS store functionality

---

## What is Gord-Overlay?
Gord-Overlay is a DHT based distribute key-value store.
Gord-Overlay will start as a REST server and your application can access data via REST.

## How is it work?
Gord-Overlay is an implementation of [DHT Chord](https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf) and simple key-value store using the DHT.
Chord protocol is an algorithm which extends consistent hashing.
Gord-Overlay, using chord protocol, allocates a key to a node in distributed nodes ring.

![chord ring](docs/architecture-1.png) 

In addition, KVS servers communicate with each other via REST to synchronize route information.
Then, servers can query via REST to resolve other servers address and communicate with these.
(REST I/F between servers is HTTP/2 but one for KVS accesses is HTTP/1.1)

## Usage
- Gord-Overlay REST server reqires a hostname and port number pair
- If you specify 127.0.0.1:26000, the server will use 127.0.0.1:26000 to communication between other nodes and use 127.0.0.1:26001 to listen local REST request
- Specified hostname and port number pair is internally used as a KVS/Proxy server identifier, so you need to specify a unique pair for each server
- Additionaly, you need to specify address of a server which is already in the DHT network to join the network except the first server
- Proxy server introduced later section is a utility to connect to the overlay network from outside of the network
  - Proxy connect to bootstrap server. currently, proxy server at ryogrid.net:9999 is hard-coded
  - But you can change it by modifying [this line](https://github.com/ryogrid/gossip-port-forward/blob/master/constants/constants.go#L3)
  - First proxy can launch with command `./third/gossip-port-forward/gossip-port-forward relay -p listenPort(required)`
  - Other proxies can be bootstrap server also. But one proxy server which has global IP is needed at least on overlay network for NAT travarsal
```
## Build
make build

## Build proxy server for gordolctl connecting to overlay network
git pull
git submodule init
git submodule update
cd third/gossip-port-forward
go build -o gossip-port-forward gossip-port-forward.go

## Start proxy server (you do this only if you launch bootstrap proxy server myself)
./third/gossip-port-forward/gossip-port-forward both -a 127.0.0.1 -f forwardAddress(required) -l listenPort(required)

## Start KVS server (internally proxy server is launched) 
./gordolctl -l hostAndPort(required) -p proxyHostAndPort(required) -n existNodeHostAndPort(optional) 
```

## Examples

1.A Start servers (after binary builds described above)
```bash
# launch 3 servers (3 shells are needed...)
cp ./third/gossip-port-forward/gossip-port-forward .

# each gordlctl internally launch gossip-port-forward process for myself
./gordolctl -l 127.0.0.1:20000
./gordolctl -l 127.0.0.1:20004 -n 127.0.0.1:20000 
./gordolctl -l 127.0.0.1:20008 -n 127.0.0.1:20004 
```

1.B Start servers (using docker-compose)
```bash
git clone https://github.com/ryogrid/gord-overlay.git
cd gord-overlay
git pull
git submodule init
git submodule update
docker-compose build && docker-compose up
```

2. Try! 
```bash
# Query
curl -X POST -H "Content-Type: application/json" -d '{"key": "gord1"}' http://localhost:20001/server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord1"}' http://localhost:20005/server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord1"}' http://localhost:20009/server.ExternalService/FindHostForKey 

curl -X POST -H "Content-Type: application/json" -d '{"key": "gord2"}' http://localhost:20001/server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord2"}' http://localhost:20005 server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord2"}' http://localhost:20009 server.ExternalService/FindHostForKey 

curl -X POST -H "Content-Type: application/json" -d '{"key": "gord"}' localhost:20001 server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord"}' localhost:20005 server.ExternalService/FindHostForKey \
&& curl -X POST -H "Content-Type: application/json" -d '{"key": "gord"}' localhost:20009 server.ExternalService/FindHostForKey

# Put and Get value
curl -X POST -H "Content-Type: application/json" -d '{"key": "hoge", "value": "foobar"}' http://localhost:20001/server.ExternalService/PutValue
curl -X POST -H "Content-Type: application/json" -d '{"key": "hoge"}' http://localhost:20005/server.ExternalService/GetValue
curl -X POST -H "Content-Type: application/json" -d '{"key": "hoge"}' http://localhost:20009/server.ExternalService/GetValue
```

## Development Memo
- [here (Japanese)](https://zenn.dev/ryogrid/scraps/42d5c81e8604fd)

## TODO (not implemented part)
- Data replication
  - Puted data is stored only one server now
- Abnormal situation handling
  - Servers may crash when a node leavs DHT network or HTTP/2 connection...
- Handling new server join and server leave after data writing started
  - When join, data delegation is needed
  - When leave, number of servers which have replica entries (key/value) same with leaved node should be keeped 
    - Data replication to additinal server must occur


