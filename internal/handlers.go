package internal

import "io"

type Handler func(w io.Writer, r *Request)
