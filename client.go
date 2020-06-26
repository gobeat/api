package api

import (
	"net/http"
	"net/url"
	"time"
)

// Client represents for Client Client
type Client interface {
	// URL sets request's URL
	URL(url *url.URL) Client

	// Method specifies request's method
	Method(method string) Client

	// Body sets request's body
	// It is used in POST, PUT, PATCH requests
	Body(body interface{}) Client

	// Header sets request's header
	Header(key string, value string) Client

	// Query sets request's query
	// It will be added in query of url
	Query(key string, value string) Client

	// Use allows to add middleware
	Use(m Middleware) Client

	// Timeout allows to set request's timeout
	Timeout(timeout time.Duration) Client

	// Send executes the request
	Send() (*http.Response, error)
}
