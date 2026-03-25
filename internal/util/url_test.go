// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	testCases := []struct {
		Name     string
		BaseUrl  string
		Path     string
		Query    map[string]string
		Expected string
	}{
		{
			Name:    "String query",
			BaseUrl: "https://example.com",
			Path:    "/api/v1/notify",
			Query: map[string]string{
				"hello": "world",
			},
			Expected: "https://example.com/api/v1/notify?hello=world",
		},
		{
			Name:    "Bool query",
			BaseUrl: "https://example.com",
			Path:    "/api/v1/notify",
			Query: map[string]string{
				"atest": fmt.Sprintf("%v", true),
			},
			Expected: "https://example.com/api/v1/notify?atest=true",
		},
	}

	for _, input := range testCases {
		t.Run(input.Name, func(t *testing.T) {
			assertions := assert.New(t)

			url, err := BuildUrl(input.BaseUrl, input.Path, input.Query)
			assertions.NoError(err)
			assertions.Equal(input.Expected, url)
		})
	}
}
