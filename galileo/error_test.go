// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package galileo

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapError(t *testing.T) {
	testCases := []struct {
		Name     string
		Err      error
		Expected bool
	}{
		{
			Name:     "non-galileo-error",
			Err:      fmt.Errorf("not an galileo error"),
			Expected: false,
		},
		{
			Name:     "galileo-error",
			Err:      wrapError(fmt.Errorf("galileo error")),
			Expected: true,
		},
	}

	for _, input := range testCases {
		t.Run(input.Name, func(t *testing.T) {
			assertions := assert.New(t)
			assertions.Equal(input.Expected, errors.Is(input.Err, ErrGalileo))
		})
	}
}
