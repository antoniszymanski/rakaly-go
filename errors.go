// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build (windows && amd64) || (darwin && arm64) || (linux && amd64)

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
