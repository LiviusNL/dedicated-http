// Copyright 2023 LiviusNL
package dedicatedhttp_test

import (
	"net/http"
	"testing"

	dedicatedhttp "github.com/liviusnl/dedicated-http"
)

func TestHTTPClient(t *testing.T) {
	dhc := dedicatedhttp.NewClient()

	if dhc == http.DefaultClient {
		t.Error("NewClient return http.DefaultClient; want new *http.Client")
	}
}

func TestHTTPClientDefaultTransport(t *testing.T) {
	dhc := dedicatedhttp.NewClient()

	if dhc.Transport == http.DefaultTransport {
		t.Error("NewClient() uses http.DefaultTransport; want new *http.Transport")
	}
}

func TestHTTPClientDedicatedTransport(t *testing.T) {
	dhc1 := dedicatedhttp.NewClient()
	dhc2 := dedicatedhttp.NewClient()

	if dhc1.Transport == dhc2.Transport {
		t.Error("NewClient() uses *Client().Transport; want new *http.Transport")
	}
}
