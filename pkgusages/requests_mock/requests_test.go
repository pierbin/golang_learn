package main

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

/**
http-mock can be used to do following things.
Mock Response — to set exactly what we want to return
Mock Request URL — a url which we want to intercept (httpmock supports wildcards as well)
Input parameters — a set of parameters which would trigger the correct url and response from previous points
*/
func TestRequests(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder(
		"GET",
		"https://baidu.com",
		httpmock.NewStringResponder(200, "resp string"),
	)

	resp, _ := http.Get("https://baidu.com")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "resp string", string(body))
}
