package http

import (
	"bufio"
	"io"
	"net/url"
	"strings"
)

type Request struct {
	Method string
	URL    url.URL
	Proto  string
	Body   io.ReadCloser
}

func NewRequest(method, urlStr string, body io.Reader) (*Request, error) {
	req := new(Request)
	req.Method = method
	req.URL.Path = urlStr
	req.Proto = "HTTP/1.1"
	req.Body = body.(io.ReadCloser)
	return req, nil
}

func ReadRequest(r io.Reader) (*Request, error) {
	data, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		return nil, err
	}
	s := strings.SplitN(data, " ", 3)
	if len(s) < 3 {
		//TODO: Bad request
		return nil, nil
	}
	//	proto := strings.Split(s[2], "\r\n")
	return NewRequest(s[0], s[1], r)
}
