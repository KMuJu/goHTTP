package goHTTP

type Code int

var (
	Ok                  Code = 200
	NotFound            Code = 400
	Found               Code = 302
	InternalServerError Code = 500
)
