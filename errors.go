//go:build (windows && amd64) || (darwin && arm64) || (linux && amd64)

/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package rakaly

type Error string

func (e Error) Error() string {
	return "rakaly: " + string(e)
}

func (e Error) Unwrap() error {
	if e == Err {
		return nil
	}
	return Err
}

const (
	Err                 = Error("error")
	ErrAlreadyPlaintext = Error("already plaintext")
	ErrCopyFailed       = Error("failed to copy data")
	ErrMetadataNotFound = Error("metadata not found")
)
