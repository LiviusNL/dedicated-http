// Copyright 2023 LiviusNL
package dedicatedhttp_test

import (
	"net/http"
	"testing"

	dedicatedhttp "github.com/liviusnl/dedicated-http"
	"github.com/stretchr/testify/assert"
)

func TestDedicatedHTTP(t *testing.T) {
	dhc := dedicatedhttp.NewClient()

	assert.NotEqual(t, http.DefaultClient, dhc,
		"NewClient returns default HTTP Client; want new dedicated HTTP client")

	assert.NotEqual(t, dedicatedhttp.NewClient(), dhc,
		"NewClient should return a new dedicated HTTP client with each call")

	assert.NotEqual(t, http.DefaultTransport, dhc.Transport,
		"NewClient returns default HTTP Transport; want new dedicated HTTP transport ")

	assert.NotEqual(t, dedicatedhttp.NewClient().Transport, dhc.Transport,
		"NewClient should return a new HTTP client with a new HTTP Transport with each call")

}
