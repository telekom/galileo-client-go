// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package options

import "time"

type ClientOptions struct {
	Debug   bool
	DryRun  bool
	BaseUrl string
	Timeout time.Duration
	Auth    *AuthOptions
}

func Client() *ClientOptions {
	opts := &ClientOptions{
		Debug:   false,
		DryRun:  true,
		BaseUrl: "http://localhost:8080",
		Auth: &AuthOptions{
			Enabled:      false,
			TokenUrl:     "https://example.com/oauth/token",
			ClientId:     "galileo",
			ClientSecret: "not-a-secret",
		},
		Timeout: 30 * time.Second,
	}
	return opts
}

func (o *ClientOptions) SetDebug(debug bool) *ClientOptions {
	o.Debug = debug
	return o
}

func (o *ClientOptions) SetDryRun(dryRun bool) *ClientOptions {
	o.DryRun = dryRun
	return o
}

func (o *ClientOptions) SetBaseUrl(baseUrl string) *ClientOptions {
	o.BaseUrl = baseUrl
	return o
}

func (o *ClientOptions) SetTimeout(timeout time.Duration) *ClientOptions {
	o.Timeout = timeout
	return o
}

func (o *ClientOptions) WithAuth(auth *AuthOptions) *ClientOptions {
	o.Auth = auth
	return o
}
