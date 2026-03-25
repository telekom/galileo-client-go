// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Dump struct {
	Request  []byte
	Response []byte
}

func NewDump() *Dump {
	return &Dump{
		Request:  make([]byte, 0),
		Response: make([]byte, 0),
	}
}

func (d *Dump) RecordRequest(req *http.Request, debug bool) {
	if debug {
		bytes, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			panic(err)
		}
		d.Request = append(d.Request, bytes...)
	}
}

func (d *Dump) RecordResponse(resp *http.Response, debug bool) {
	if debug {
		bytes, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		d.Response = append(d.Response, bytes...)
	}
}

func (d *Dump) Print() {
	builder := new(strings.Builder)
	_, _ = builder.WriteString("### Request\n")
	_, _ = builder.Write(d.Request)

	_, _ = builder.WriteString("\n\n### Response\n")
	_, _ = builder.Write(d.Response)

	fmt.Printf("%s\n", builder.String())
}
