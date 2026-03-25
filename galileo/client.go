// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package galileo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/telekom/galileo-client-go/galileo/response"
	"github.com/telekom/galileo-client-go/internal/util"
	"github.com/telekom/galileo-client-go/options"
)

type Client struct {
	client *http.Client
	opts   *options.ClientOptions
}

func NewClient(opts ...*options.ClientOptions) *Client {
	clientOpts := util.IfThenElse(len(opts) > 0, opts[0], options.Client())
	client := &http.Client{
		Timeout: clientOpts.Timeout,
	}

	return &Client{
		client: util.IfThenElse(clientOpts.Auth.Enabled, clientOpts.Auth.Authorize(client), client),
		opts:   clientOpts,
	}
}

func (c *Client) Notify(resolver string, opts *options.NotifyOptions) (*response.NotifyResponse, error) {
	url, err := util.BuildUrl(c.opts.BaseUrl, "/api/v1/notify", map[string]string{
		"resolver": strings.ToUpper(resolver),
		"dryrun":   fmt.Sprintf("%v", c.opts.DryRun),
	})
	if err != nil {
		return nil, wrapError(err)
	}

	dump := util.NewDump()
	req, err := c.createNotifyRequest(url, opts)
	if err != nil {
		return nil, wrapError(err)
	}
	dump.RecordRequest(req, c.opts.Debug)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, wrapError(err)
	}
	defer res.Body.Close()
	dump.RecordResponse(res, c.opts.Debug)

	if res.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("received unexpected status code: %d", res.StatusCode)
		return nil, wrapError(err)
	}

	if c.opts.Debug {
		dump.Print()
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, wrapError(err)
	}

	var responseBody response.NotifyResponse
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return nil, wrapError(err)
	}

	return &responseBody, nil
}

func (c *Client) createNotifyRequest(url string, opts *options.NotifyOptions) (*http.Request, error) {
	reqBody, err := opts.JsonBytes()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) GetClient() *http.Client {
	return c.client
}
