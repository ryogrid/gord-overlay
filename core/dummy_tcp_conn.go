package core

import (
	"net"
)

type DummyTCPConn struct {
	net.Conn
	dummyRemoteAddr net.Addr
}

func NewDummyTCPConn(orgTCPConn *net.Conn, dummyRemoteAddr net.Addr) *DummyTCPConn {
	return &DummyTCPConn{*orgTCPConn, dummyRemoteAddr}
}

func (dtc *DummyTCPConn) RemoteAddr() net.Addr {
	return dtc.dummyRemoteAddr
}
