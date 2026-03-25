// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

//go:build exclude

package main

import (
	"fmt"
	"time"

	"github.com/telekom/galileo-client-go/galileo"
	"github.com/telekom/galileo-client-go/options"
)

func main() {
	// Begin by creating the client configuration
	clientOpts := options.Client().
		// Whether to dump requests and responses to the console
		SetDebug(false).
		// Whether the requests should be sent in dry-run mode
		SetDryRun(false).
		// The base url of the galileo service
		SetBaseUrl("https://galileo.example.com/").
		// Http timeout
		SetTimeout(30 * time.Second).
		// Apply authentication options
		WithAuth(
			options.Auth().
				// Whether requests should be authenticated or not
				SetEnabled(true).
				// The issuers token endpoint
				SetIssuer("https://idp.example.com/token").
				// The client id as required by the oauth client-credentials flow
				SetClientId("myclient").
				// The client secret as required by the oauth client-credentials flow
				SetClientSecret("mysecret"),
		)

	// Create the client
	client := galileo.NewClient(clientOpts)

	// Prepare the configuration for the notification
	notifyOpts := options.Notify().
		// We plan to use the email resolver, so we pass an email address as the parameter
		SetParameters([]string{"john.doe@example.com"}).
		// The sender email address
		SetSender("jane.doe@example.com").
		// The sender name
		SetSenderName("Jane Doe").
		// The email address the summary should be sent to after processing all addresses
		SetSummaryRecipient("jane.doe@example.com").
		// The file name of the template that should be used
		SetTemplate("mytemplate.tmpl").
		// The subject of the email
		SetSubject("Hello World").
		// The data used for rendering the template
		SetData(map[string]any{
			"foo": "bar",
		})

	res, err := client.Notify("EMAIL", notifyOpts)
	if err != nil {
		panic("could not send emails")
	}
	fmt.Printf("Response message: %s\n", res.Message)
}
