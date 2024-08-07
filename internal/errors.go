package internal

type methodNotSupported struct{}
type errorInHeader struct{}
type notFound struct{}

func (err methodNotSupported) Error() string { return "Method not supported" }
func (err errorInHeader) Error() string      { return "Error in header" }
func (err notFound) Error() string           { return "Not Found" }
