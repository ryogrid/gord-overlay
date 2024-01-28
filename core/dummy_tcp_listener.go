package core

import (
	"fmt"
	"github.com/ryogrid/gossip-overlay/gossip"
	"io"
	"math"
	"net"
	"net/netip"
)

type DummyTCPListener struct {
	net.TCPListener
}

func NewDummyTCPListener(listenAddrStr string) net.Listener {
	dummyListener, err := net.ListenTCP("tcp", net.TCPAddrFromAddrPort(netip.MustParseAddrPort(listenAddrStr)))
	if err != nil {
		panic(err)
	}

	return &DummyTCPListener{*dummyListener}
}

// Accept waits for and returns the next connection to the listener.
func (dtl *DummyTCPListener) Accept() (net.Conn, error) {
	conn, err := dtl.TCPListener.Accept()
	if err != nil {
		fmt.Println("DummyTCPListener::Accept failed", err)
	}

	// read remote node address at start of stream wrote by proxy
	lenBuf := make([]byte, 1)
	// read 1 byte
	n, err := io.ReadFull(conn, lenBuf)
	if err != nil || n != 1 {
		fmt.Println("DummyTCPListener::Accept failed (reading addres len)", err)
	}
	addrStrLen := int(lenBuf[0])
	addrBuf := make([]byte, addrStrLen)
	// read addrStrLen bytes
	n, err = io.ReadFull(conn, addrBuf)
	if err != nil || n != addrStrLen {
		fmt.Println("DummyTCPListener::Accept failed (reading address)", err)
	}
	remoteAddrStr := string(addrBuf)

	// set remote node address to conn (not address of proxy)
	retConn := &DummyTCPConn{conn, &gossip.PeerAddress{
		PeerName: math.MaxUint64,
		PeerHost: &remoteAddrStr,
	}}
	return retConn, err
}
