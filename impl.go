package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type defaultClient struct {
	r *http.Request
	e error
	t time.Duration
}

func newClient(method string, rawurl string) Client {
	c := &defaultClient{
		r: new(http.Request),
	}

	c.r.Method = method
	u, e := url.Parse(rawurl)
	if e != nil {
		c.e = e
	} else {
		c.r.URL = u
	}
	c.r.Header = make(http.Header)

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
	var b string
	if bytes, ok := body.([]byte); ok {
		b = string(bytes)
	} else if str, ok := body.(string); ok {
		b = str
	} else {
		bytes, e := json.Marshal(body)
		if e != nil {
			c.e = e
			return c
		}

		return c.Body(bytes)
	}

	c.r.Body = ioutil.NopCloser(strings.NewReader(b))
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
