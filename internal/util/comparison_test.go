// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfThenElse(t *testing.T) {
	testCases := []struct {
		Name       string
		Condition  bool
		TrueValue  any
		FalseValue any
		Expected   any
	}{
		{
			Name:       "True",
			Condition:  true,
			TrueValue:  "foo",
			FalseValue: "bar",
			Expected:   "foo",
		},
		{
			Name:       "False",
			Condition:  false,
			TrueValue:  "foo",
			FalseValue: "bar",
			Expected:   "bar",
		},
	}

	for _, input := range testCases {
		t.Run(input.Name, func(t *testing.T) {
			assertions := assert.New(t)
			result := IfThenElse(input.Condition, input.TrueValue, input.FalseValue)
			assertions.Equal(input.Expected, result)
		})
	}
}
