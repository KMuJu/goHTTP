package internal

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
