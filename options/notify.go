// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package options

import "encoding/json"

type NotifyOptions struct {
	Parameters       []string       `json:"params"`
	Sender           string         `json:"sender"`
	SenderName       string         `json:"sender_name"`
	SummaryRecipient string         `json:"summary_recipient"`
	Template         string         `json:"template"`
	Subject          string         `json:"subject"`
	Data             map[string]any `json:"data"`
}

func Notify() *NotifyOptions {
	opts := &NotifyOptions{
		Sender:     "galileo-client@example.com",
		SenderName: "Galileo",
	}
	return opts
}

func (o *NotifyOptions) SetParameters(parameters []string) *NotifyOptions {
	o.Parameters = parameters
	return o
}

func (o *NotifyOptions) SetSender(sender string) *NotifyOptions {
	o.Sender = sender
	return o
}

func (o *NotifyOptions) SetSenderName(name string) *NotifyOptions {
	o.SenderName = name
	return o
}

func (o *NotifyOptions) SetSummaryRecipient(recipient string) *NotifyOptions {
	o.SummaryRecipient = recipient
	return o
}

func (o *NotifyOptions) SetTemplate(template string) *NotifyOptions {
	o.Template = template
	return o
}

func (o *NotifyOptions) SetSubject(subject string) *NotifyOptions {
	o.Subject = subject
	return o
}

func (o *NotifyOptions) SetData(data map[string]any) *NotifyOptions {
	o.Data = data
	return o
}

func (o *NotifyOptions) JsonBytes() ([]byte, error) {
	return json.Marshal(o)
}
