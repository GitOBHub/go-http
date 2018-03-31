package http

import (
	"log"
	"net"
)

type Server struct {
	Addr string
}

func (srv *Server) Listen(proto, addr string) (net.Listener, error) {
	if addr == "" {
		srv.Addr = addr
	}
	return net.Listen(proto, srv.Addr)
}

func (srv *Server) Serve(ln net.Listener) error {
	c, err := ln.Accept()
	if err != nil {
		return error
	}
	conn := srv.newConn(c)
	go conn.Serve()
}

func (srv *Server) newConn(
