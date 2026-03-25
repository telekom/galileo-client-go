// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package util

import "net/url"

// BuildUrl concatenates the given base-url and path before adding the given query parameters.
// The given path MUST begin with a leading slash!
func BuildUrl(baseUrl string, path string, query map[string]string) (string, error) {
	fullUrl, err := url.Parse(baseUrl + path)
	if err != nil {
		return "", err
	}

	queryParams := fullUrl.Query()
	for k, v := range query {
		queryParams.Set(k, v)
	}
	fullUrl.RawQuery = queryParams.Encode()

	return fullUrl.String(), nil
}
