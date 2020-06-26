# API

## Getting Started

- `Gopkg.toml`

```
[[constraint]]
  name = "github.com/gobeat/api"
  version = "~1.0.x"
```

- Make request

```go
import (
  "github.com/gobeat/api"
)

// ...

prefix := func(r *http.Request) {
  r.URL.Scheme = "http"
  r.URL.Host = "localhost:3030"
}
response, err := api.GET("/path").
    Header("Authorization", "Bearer token"). // set header
    Body("this is a body"). // set body
    Query("foo", "baz"). // set query
    Use(prefix). // use prefix middleware
    Send()
```