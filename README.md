# Go HTTP Client with dedicated HTTP Transport

The Go `http` standard library package implements one shared HTTP Transport `http.DefaultTransport`.
When not explicitly overridden all HTTP clients use this Transport.

This package implements a dedicated `http.Transport`
using the same default values as `http.DefaultTransport`.