// Copyright 2023 LiviusNL
package dedicatedhttp_test

import (
	"net/http"
	"testing"

	dedicatedhttp "github.com/liviusnl/dedicated-http"
	"github.com/stretchr/testify/assert"
)

func TestDedicatedHTTP(t *testing.T) {
	dhc1 := dedicatedhttp.NewClient()
	dhc2 := dedicatedhttp.NewClient()

	assert.NotEqual(t, dhc1, http.DefaultClient,
		"NewClient returns default HTTP Client; want new dedicated HTTP client")

	assert.NotEqual(t, dhc1, dhc2,
		"NewClient should return a new dedicated HTTP client with each call")

	assert.NotEqual(t, dhc1.Transport, http.DefaultTransport,
		"NewClient returns default HTTP Transport; want new dedicated HTTP transport ")

	assert.NotEqual(t, dhc1.Transport, dhc2.Transport,
		"NewClient should return a new HTTP client with a new HTTP Transport with each call")

}
