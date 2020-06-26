package api_test

import (
	"io/ioutil"
	"net/http"

	"github.com/gobeat/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {
	prefix := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = "localhost:3030"
	}

	It("should execute GET request", func() {
		r, e := api.GET("/users").
			Use(prefix).
			Header("content-type", "application/json").
			Send()
		Expect(e).To(BeNil())
		Expect(r.StatusCode).To(Equal(http.StatusOK))
		Expect(r.Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		body, e := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		Expect(e).To(BeNil())
		Expect(string(body)).To(Equal("[{\"id\":1,\"name\":\"John\"},{\"id\":2,\"name\":\"Marry\"}]"))
	})

	It("should execute GET request with query", func() {
		r, e := api.GET("/users").
			Use(prefix).
			Header("content-type", "application/json").
			Query("gender", "male").
			Send()
		Expect(e).To(BeNil())
		Expect(r.StatusCode).To(Equal(http.StatusOK))
		Expect(r.Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		body, e := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		Expect(e).To(BeNil())
		Expect(string(body)).To(Equal("[{\"id\":1,\"name\":\"John\"}]"))
	})

	It("should execute POST request", func() {
		r, e := api.POST("/users").
			Use(prefix).
			Body(map[string]string{
				"name": "John",
			}).
			Header("content-type", "application/json").
			Send()
		Expect(e).To(BeNil())
		Expect(r.StatusCode).To(Equal(http.StatusCreated))
		Expect(r.Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		body, e := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		Expect(e).To(BeNil())
		Expect(string(body)).To(Equal("{\"id\":1,\"name\":\"John\"}"))
	})
})
