// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package galileo

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/galileo-client-go/galileo/response"
	"github.com/telekom/galileo-client-go/internal/util"
	"github.com/telekom/galileo-client-go/options"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestClient_Notify(t *testing.T) {
	opts := options.Client().SetBaseUrl("http://mock.local").SetDebug(true).SetDryRun(false)
	client := NewClient(opts)

	activateMock(client)

	testCases := []struct {
		Name    string
		Options *options.NotifyOptions
	}{
		{
			Name: "ALL_USERS",
			Options: options.Notify().
				SetSender("john.doe@example.com").
				SetSenderName("John Doe").
				SetSummaryRecipient("john.doe@example.com").
				SetTemplate("mytemplate.tmpl").
				SetSubject("Hello World").
				SetData(map[string]any{
					"foo":  "bar",
					"fizz": "buzz",
				}),
		},
		{
			Name: "EMAIL",
			Options: options.Notify().
				SetParameters([]string{"john.doe@example.com", "jane.doe@example.com"}).
				SetSender("john.doe@example.com").
				SetSenderName("John Doe").
				SetSummaryRecipient("john.doe@example.com").
				SetTemplate("mytemplate.tmpl").
				SetSubject("Hello World").
				SetData(map[string]any{
					"foo":  "bar",
					"fizz": "buzz",
				}),
		},
		{
			Name: "TEAM",
			Options: options.Notify().
				SetParameters([]string{"team1", "team2"}).
				SetSender("john.doe@example.com").
				SetSenderName("John Doe").
				SetSummaryRecipient("john.doe@example.com").
				SetTemplate("mytemplate.tmpl").
				SetSubject("Hello World").
				SetData(map[string]any{
					"foo":  "bar",
					"fizz": "buzz",
				}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assertions := assert.New(t)

			res, err := client.Notify(testCase.Name, testCase.Options)
			assertions.NoError(err)
			assertions.True(len(res.Message) > 0, "missing message in response")
			assertions.Equal(testCase.Name, res.Details.Resolver, "missing resolver in response")
			assertions.Equal(testCase.Options.Parameters, res.Details.Params, "missing parameters in response")
			assertions.Equal(testCase.Options.Template, res.Details.Template, "missing template in response")
		})
	}
}

func activateMock(client *Client) {
	httpmock.ActivateNonDefault(client.client)

	httpmock.RegisterResponder(http.MethodPost, "http://mock.local/api/v1/notify", func(request *http.Request) (*http.Response, error) {
		resolver := strings.ToLower(request.URL.Query().Get("resolver"))
		switch resolver {

		case "all_users", "email", "team":
			var reqBody struct {
				Params   []string `json:"params"`
				Template string   `json:"template"`
			}

			bodyBytes, _ := io.ReadAll(request.Body)
			_ = json.Unmarshal(bodyBytes, &reqBody)

			resBody := response.NotifyResponse{
				Message: util.IfThenElse(!client.opts.DryRun, "Emails are being sent asynchronously", "Dryrun enabled, email will be sent to the sender only"),
				Details: struct {
					Resolver string   `json:"resolver"`
					Params   []string `json:"params"`
					Template string   `json:"template"`
				}{Resolver: strings.ToUpper(resolver), Params: reqBody.Params, Template: reqBody.Template},
			}
			return httpmock.NewJsonResponse(http.StatusAccepted, resBody)

		default:
			return httpmock.NewJsonResponse(http.StatusBadRequest, map[string]any{
				"message": "Unknown resolver, must be one of ALL_USERS, EMAIL, TEAMS",
			})
		}
	})
}
