package internal

type method string

type message struct {
	raw []byte
}

type Header map[string][]string

type Request struct {
	Method  string
	URL     string
	Version string
	Header  Header
	Body    []byte
}

type RequestWriter struct {
	bytes []byte
}
