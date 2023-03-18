package service

type Context interface {
	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error
	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error
	// Param returns path parameter by name.
	Param(name string) string
}
