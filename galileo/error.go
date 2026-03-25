// Copyright 2026 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package galileo

import "errors"

var ErrGalileo = errors.New(`galileo: `)

func wrapError(err error) error {
	return errors.Join(ErrGalileo, err)
}
