// Copyright 2023 LiviusNL
package dedicatedhttp

import "net/http"

// Client returns a new http.Client with a dedicated http.Transport
func NewClient() *http.Client {
	return &http.Client{
		Transport: Transport(),
	}
}
