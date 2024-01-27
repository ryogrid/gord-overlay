package core

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/gossip"
	"math"
	"net"
	"net/netip"
)

type DummyTCPListener struct {
	net.TCPListener
	dummyRemoteHost net.Addr
}

func NewDummyTCPListener(listenAddrStr string, dummyRemoteHost string) net.Listener {
	dummyListener, err := net.ListenTCP("tcp", net.TCPAddrFromAddrPort(netip.MustParseAddrPort(listenAddrStr)))
	if err != nil {
		panic(err)
	}

	return &DummyTCPListener{*dummyListener, &gossip.PeerAddress{
		PeerName: math.MaxUint64,
		PeerHost: &dummyRemoteHost,
	}}
}

// Accept waits for and returns the next connection to the listener.
func (dtl *DummyTCPListener) Accept() (net.Conn, error) {
	conn, err := dtl.TCPListener.Accept()
	if err != nil {
		fmt.Println("DummyTCPListener::Accept failed", err)
	}

	retConn := &DummyTCPConn{conn, dtl.dummyRemoteHost}
	return retConn, err
}
