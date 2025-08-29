// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build (windows && amd64) || (darwin && arm64) || (linux && amd64)

package rakaly

import "unsafe"

func bytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func resize(b []byte, length int) []byte {
	if length <= cap(b) {
		return b[:length]
	}
	return makeNoZero(length)
}

//go:linkname makeNoZero internal/bytealg.MakeNoZero
func makeNoZero(length int) []byte
