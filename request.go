package http

type Request struct {
	Method string
	Url    url.URL
	Proto  string
	Body   io.ReadCloser
}
