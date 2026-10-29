// Copyright 2026 Umut Alisoglu
//
// SPDX-License-Identifier: Apache-2.0

package galileo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/telekom/galileo-client-go/options"
)

func TestNewClientDefaultOptions(t *testing.T) {
	client := NewClient()
	assert.Equal(t, options.Client(), client.opts)
	assert.Equal(t, client.opts.Timeout, client.GetClient().Timeout)
	assert.True(t, client.opts.DryRun)
}

func TestNewClientExplicitOptions(t *testing.T) {
	first := options.Client().SetBaseUrl("https://first.example").SetTimeout(time.Second)
	second := options.Client().SetBaseUrl("https://second.example").SetTimeout(2 * time.Second)

	for _, opts := range [][]*options.ClientOptions{{first}, {first, second}} {
		client := NewClient(opts...)
		assert.Same(t, first, client.opts)
		assert.Equal(t, time.Second, client.GetClient().Timeout)
	}
}
