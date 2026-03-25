// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package options

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

type AuthOptions struct {
	Enabled      bool
	TokenUrl     string
	ClientId     string
	ClientSecret string
}

func Auth() *AuthOptions {
	return &AuthOptions{}
}

func (o *AuthOptions) SetEnabled(enabled bool) *AuthOptions {
	o.Enabled = enabled
	return o
}

func (o *AuthOptions) SetIssuer(issuerUrl string) *AuthOptions {
	o.TokenUrl = issuerUrl
	return o
}

func (o *AuthOptions) SetClientId(clientId string) *AuthOptions {
	o.ClientId = clientId
	return o
}

func (o *AuthOptions) SetClientSecret(clientSecret string) *AuthOptions {
	o.ClientSecret = clientSecret
	return o
}

func (o *AuthOptions) Authorize(client *http.Client) *http.Client {
	credentials := clientcredentials.Config{
		ClientID:     o.ClientId,
		ClientSecret: o.ClientSecret,
		TokenURL:     o.TokenUrl,
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, client)
	oauthClient := credentials.Client(ctx)
	return oauthClient
}
