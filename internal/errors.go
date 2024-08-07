package internal

import "errors"

var (
	methodNotSupported = errors.New("Method not supported")
	errorInHeader      = errors.New("Error in header")
	notFound           = errors.New("Not Found")
)
