package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type defaultClient struct {
	r *http.Request
	e error
	t time.Duration
}

func newClient(method string, url string) Client {
	req, e := http.NewRequest(method, url, nil)
	if e != nil {
		panic(e)
	}

	c := &defaultClient{
		r: req,
	}

	return c
}

func (c *defaultClient) URL(url *url.URL) Client {
	c.r.URL = url
	return c
}

func (c *defaultClient) Method(method string) Client {
	c.r.Method = method
	return c
}

func (c *defaultClient) Body(body interface{}) Client {
	var buf []byte
	if bytes, ok := body.([]byte); ok {
		buf = bytes
	} else if str, ok := body.(string); ok {
		buf = []byte(str)
	} else {
		bytes, e := json.Marshal(body)
		if e != nil {
			c.e = e
			return c
		}
		buf = bytes
	}

	c.r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return c
}

func (c *defaultClient) Header(key string, value string) Client {
	c.r.Header.Set(key, value)

	return c
}

func (c *defaultClient) Query(key string, value string) Client {
	q := c.r.URL.Query()
	q.Set(key, value)
	c.r.URL.RawQuery = q.Encode()

	return c
}

func (c *defaultClient) Use(m Middleware) Client {
	m(c.r)

	return c
}

func (c *defaultClient) Timeout(timeout time.Duration) Client {
	c.t = timeout
	return c
}

func (c *defaultClient) Send() (*http.Response, error) {
	if c.e != nil {
		return nil, c.e
	}

	return (&http.Client{
		Timeout: c.t,
	}).Do(c.r)
}
