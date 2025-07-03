/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package cgo

/*
#cgo CPPFLAGS: -O2
#cgo LDFLAGS: -L${SRCDIR}/resources -lrakaly -Wl,-rpath,$ORIGIN

#include <stdbool.h>
#include "resources/rakaly.h"
*/
import "C"
import "unsafe"

type (
	MeltedBuffer       = *C.MeltedBuffer
	MeltedBufferResult = *C.MeltedBufferResult
	PdsError           = *C.PdsError
	PdsFile            = *C.PdsFile
	PdsFileResult      = *C.PdsFileResult
	PdsMeta            = *C.PdsMeta
)

func Rakaly_free_melt(p0 MeltedBuffer) {
	C.rakaly_free_melt(p0)
}
func Rakaly_melt_data_length(p0 MeltedBuffer) uint {
	return uint(C.rakaly_melt_data_length(p0))
}
func Rakaly_melt_is_verbatim(p0 MeltedBuffer) bool {
	return bool(C.rakaly_melt_is_verbatim(p0))
}
func Rakaly_melt_binary_unknown_tokens(p0 MeltedBuffer) bool {
	return bool(C.rakaly_melt_binary_unknown_tokens(p0))
}
func Rakaly_melt_write_data(p0 MeltedBuffer, p1 *byte, p2 uint) uint {
	return uint(C.rakaly_melt_write_data(p0, (*C.char)(unsafe.Pointer(p1)), C.size_t(p2)))
}
func Rakaly_file_error(p0 PdsFileResult) PdsError {
	return C.rakaly_file_error(p0)
}
func Rakaly_error_length(p0 PdsError) int {
	return int(C.rakaly_error_length(p0))
}
func Rakaly_error_write_data(p0 PdsError, p1 *byte, p2 int) int {
	return int(C.rakaly_error_write_data(p0, (*C.char)(unsafe.Pointer(p1)), C.int(p2)))
}
func Rakaly_free_error(p0 PdsError) {
	C.rakaly_free_error(p0)
}
func Rakaly_free_file(p0 PdsFile) {
	C.rakaly_free_file(p0)
}
func Rakaly_file_value(p0 PdsFileResult) PdsFile {
	return C.rakaly_file_value(p0)
}
func Rakaly_file_meta(p0 PdsFile) PdsMeta {
	return C.rakaly_file_meta(p0)
}
func Rakaly_file_meta_melt(p0 PdsMeta) MeltedBufferResult {
	return C.rakaly_file_meta_melt(p0)
}
func Rakaly_file_melt(p0 PdsFile) MeltedBufferResult {
	return C.rakaly_file_melt(p0)
}
func Rakaly_file_is_binary(p0 PdsFile) bool {
	return bool(C.rakaly_file_is_binary(p0))
}
func Rakaly_melt_error(p0 MeltedBufferResult) PdsError {
	return C.rakaly_melt_error(p0)
}
func Rakaly_melt_value(p0 MeltedBufferResult) MeltedBuffer {
	return C.rakaly_melt_value(p0)
}
func Rakaly_eu4_file(p0 *byte, p1 uint) PdsFileResult {
	return C.rakaly_eu4_file((*C.char)(unsafe.Pointer(p0)), C.size_t(p1))
}
func Rakaly_ck3_file(p0 *byte, p1 uint) PdsFileResult {
	return C.rakaly_ck3_file((*C.char)(unsafe.Pointer(p0)), C.size_t(p1))
}
func Rakaly_imperator_file(p0 *byte, p1 uint) PdsFileResult {
	return C.rakaly_imperator_file((*C.char)(unsafe.Pointer(p0)), C.size_t(p1))
}
func Rakaly_hoi4_file(p0 *byte, p1 uint) PdsFileResult {
	return C.rakaly_hoi4_file((*C.char)(unsafe.Pointer(p0)), C.size_t(p1))
}
func Rakaly_vic3_file(p0 *byte, p1 uint) PdsFileResult {
	return C.rakaly_vic3_file((*C.char)(unsafe.Pointer(p0)), C.size_t(p1))
}
