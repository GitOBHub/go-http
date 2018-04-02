package http

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Addr string
	Handler
}

type Conn struct {
	net.Conn
	Handler
	WriteFirst bool
}

func (srv *Server) ListenAndServe() error {
	ln, err := srv.Listen()
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}

func (srv *Server) Listen() (net.Listener, error) {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	return net.Listen("tcp", addr)
}

func (srv *Server) Serve(ln net.Listener) error {
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		conn := srv.newConn(c)
		go conn.serve()
	}
	//TODO
	return nil
}

func (srv *Server) newConn(c net.Conn) Conn {
	var conn Conn
	conn.Conn = c
	conn.Handler = srv.Handler
	conn.WriteFirst = true
	return conn
}

func (c Conn) serve() {
	defer c.Close()
	for {
		req, err := ReadRequest(c.Conn)
		if err != nil {
			return
		}
		handler := c.Handler
		if handler == nil {
			handler = defaultServeMux
		}
		handler.ServeHTTP(c, req)
		//TODO
		break
	}
}

func (c Conn) Header() Header {
	header := make(Header)
	return header
}

func (c Conn) WriteHeader(status int) {
	var s string
	switch status {
	case StatusOK:
		s = "OK"
	case StatusNotFound:
		s = "NotFound"
	default: //TODO
	}
	fmt.Fprintf(c.Conn, "HTTP/1.1 %d %s\r\n\r\n", status, s)
	//	header := c.Header()
	//	header.Write(c.Conn)
}

func (c Conn) Write(data []byte) (int, error) {
	log.Printf("%s", string(data))
	if c.WriteFirst {
		c.WriteHeader(StatusOK)
		log.Print("Conn.Write: writeheader ok")
		c.WriteFirst = false
	}
	data = append(data, []byte("\r\n")...)
	return c.Conn.Write(data)
}
