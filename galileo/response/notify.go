// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package response

type NotifyResponse struct {
	Message string `json:"message"`
	Details struct {
		Resolver string   `json:"resolver"`
		Params   []string `json:"params"`
		Template string   `json:"template"`
	} `json:"details"`
}
