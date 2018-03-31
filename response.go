package http

import (
	"net"
)

type ResponseWriter interface {
	Header() Header
	WriteHeader(int)
	Write([]byte) (int, error)
}

type responseConn struct {
	net.Conn
	writeFirst bool
}

func (c *responseConn) Header() Header {
	header := make(Header)
	header.Add("Status-Line", "HTTP/1.1")
	header.Add("Status-Line", "200")
	header.Add("Status-Line", "OK")
	return header
}

func (c *responseConn) WriteHeader(int) {
	header := c.Header()
	header.Write(c.Conn)
}

func (c *responseConn) Write(data []byte) (int, error) {
	if c.writeFirst {
		c.WriteHeader(StatusOK)
		c.writeFirst = false
	}
	data = append(data, []byte("\r\n")...)
	return c.Conn.Write(data)
}
