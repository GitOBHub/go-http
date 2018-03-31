package http

import (
	"errors"
	"io"
	"strings"
)

type Header map[string][]string

func (h Header) Write(w io.Writer) error {
	statuses, ok := h["Status-Line"]
	if !ok {
		return errors.New("no status line")
	}
	resp := strings.Join(statuses, " ")
	resp += "\r\n\r\n"
	_, err := io.WriteString(w, resp)
	return err
}

func (h Header) Add(key, value string) {
	v, ok := h[key]
	if !ok {
		h[key] = []string{value}
		return
	}
	v = append(v, value)
	h[key] = v
}
