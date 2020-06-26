package api

import "net/http"

// GET returns GET client
func GET(url string) Client {
	return newClient(http.MethodGet, url)
}

// POST returns POST client
func POST(url string) Client {
	return newClient(http.MethodPost, url)
}

// PUT returns PUT client
func PUT(url string) Client {
	return newClient(http.MethodPut, url)
}

// PATCH returns PATCH client
func PATCH(url string) Client {
	return newClient(http.MethodPatch, url)
}

// DELETE returns DELETE client
func DELETE(url string) Client {
	return newClient(http.MethodDelete, url)
}

// HEAD returns HEAD client
func HEAD(url string) Client {
	return newClient(http.MethodHead, url)
}

// OPTIONS returns OPTIONS client
func OPTIONS(url string) Client {
	return newClient(http.MethodOptions, url)
}
