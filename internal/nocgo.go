// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build !cgo

package internal

import "github.com/antoniszymanski/rakaly-go/internal/nocgo"

type (
	MeltedBuffer       = nocgo.MeltedBuffer
	MeltedBufferResult = nocgo.MeltedBufferResult
	PdsError           = nocgo.PdsError
	PdsFile            = nocgo.PdsFile
	PdsFileResult      = nocgo.PdsFileResult
	PdsMeta            = nocgo.PdsMeta
)

func Rakaly_free_melt(p0 MeltedBuffer) {
	nocgo.Rakaly_free_melt(p0)
}
func Rakaly_melt_data_length(p0 MeltedBuffer) uint {
	return nocgo.Rakaly_melt_data_length(p0)
}
func Rakaly_melt_is_verbatim(p0 MeltedBuffer) bool {
	return nocgo.Rakaly_melt_is_verbatim(p0)
}
func Rakaly_melt_binary_unknown_tokens(p0 MeltedBuffer) bool {
	return nocgo.Rakaly_melt_binary_unknown_tokens(p0)
}
func Rakaly_melt_write_data(p0 MeltedBuffer, p1 *byte, p2 uint) uint {
	return nocgo.Rakaly_melt_write_data(p0, p1, p2)
}
func Rakaly_file_error(p0 PdsFileResult) PdsError {
	return nocgo.Rakaly_file_error(p0)
}
func Rakaly_error_length(p0 PdsError) int {
	return nocgo.Rakaly_error_length(p0)
}
func Rakaly_error_write_data(p0 PdsError, p1 *byte, p2 int) int {
	return nocgo.Rakaly_error_write_data(p0, p1, p2)
}
func Rakaly_free_error(p0 PdsError) {
	nocgo.Rakaly_free_error(p0)
}
func Rakaly_free_file(p0 PdsFile) {
	nocgo.Rakaly_free_file(p0)
}
func Rakaly_file_value(p0 PdsFileResult) PdsFile {
	return nocgo.Rakaly_file_value(p0)
}
func Rakaly_file_meta(p0 PdsFile) PdsMeta {
	return nocgo.Rakaly_file_meta(p0)
}
func Rakaly_file_meta_melt(p0 PdsMeta) MeltedBufferResult {
	return nocgo.Rakaly_file_meta_melt(p0)
}
func Rakaly_file_melt(p0 PdsFile) MeltedBufferResult {
	return nocgo.Rakaly_file_melt(p0)
}
func Rakaly_file_is_binary(p0 PdsFile) bool {
	return nocgo.Rakaly_file_is_binary(p0)
}
func Rakaly_melt_error(p0 MeltedBufferResult) PdsError {
	return nocgo.Rakaly_melt_error(p0)
}
func Rakaly_melt_value(p0 MeltedBufferResult) MeltedBuffer {
	return nocgo.Rakaly_melt_value(p0)
}
func Rakaly_eu4_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_eu4_file(p0, p1)
}
func Rakaly_ck3_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_ck3_file(p0, p1)
}
func Rakaly_imperator_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_imperator_file(p0, p1)
}
func Rakaly_hoi4_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_hoi4_file(p0, p1)
}
func Rakaly_vic3_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_vic3_file(p0, p1)
}
func Rakaly_eu5_file(p0 *byte, p1 uint) PdsFileResult {
	return nocgo.Rakaly_eu5_file(p0, p1)
}
