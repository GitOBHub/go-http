package http

import ()

type ResponseWriter interface {
	Header() Header
	WriteHeader(int)
	Write([]byte) (int, error)
}
