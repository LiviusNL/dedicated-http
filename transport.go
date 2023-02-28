// Copyright 2023 LiviusNL
package dedicatedhttp

import (
	"net"
	"net/http"
	"time"
)

// NewTransport returns a new, dedicated http.Transport
// with the same default values as http.DefaultTransport
//
// The default values are default for GO 1.20. See:
// https://pkg.go.dev/net/http#RoundTripper
func NewTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
