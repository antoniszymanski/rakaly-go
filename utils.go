//go:build (windows && amd64) || (darwin && arm64) || (linux && amd64)

/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package rakaly

import "unsafe"

func bytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func resize(b []byte, length int) []byte {
	if length <= cap(b) {
		return b[:length]
	}
	return append(b[:cap(b)], make([]byte, length-cap(b))...)
}
