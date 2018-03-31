package http

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

var handler Handler

const (
	StatusOK int = 200
)

func ListenAndServe(addr string, h Handler) error {
	srv := Server{Addr: ":http"}
	ln, err := srv.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	if err := srv.Serve(ln); err != nil {
		return err
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF {
		return
	}
	log.Printf("New request: %s", data)
	s := strings.SplitN(data, " ", 3)
	proto := strings.Split(s[2], "\r\n")
	if len(s) < 3 {
		io.WriteString(conn, "Bad request\r\n")
		return
	}
	req := new(Request)
	req.Method, req.Url.Path, req.Proto = s[0], s[1], proto[0]
	//req := &Request{Method: s[0], Url.Path: s[1], Proto: proto[0]}
	req.Body = conn
	respConn := &responseConn{conn, true}
	switch req.Method {
	case "GET":
		handler.ServeHTTP(respConn, req)
	default:
		io.WriteString(conn, "Bad request\r\n")
	}
}
