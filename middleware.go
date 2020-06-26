package api

import "net/http"

// Middleware represents Client's middleware
type Middleware func(r *http.Request)
