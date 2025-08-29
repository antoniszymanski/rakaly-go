// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build (windows && amd64) || (darwin && arm64) || (linux && amd64)

package rakaly

import (
	"unsafe"

	. "github.com/antoniszymanski/rakaly-go/internal"
)

func unwrapError(err PdsError) error {
	if err == nil {
		return nil
	}

	length := Rakaly_error_length(err)
	data := makeNoZero(length)
	Rakaly_error_write_data(err, unsafe.SliceData(data), length)
	Rakaly_free_error(err)
	return Error(bytesToString(data))
}

type MeltedOutput struct {
	melt MeltedBuffer
}

func (m MeltedOutput) WriteData(buf []byte) ([]byte, error) {
	if Rakaly_melt_is_verbatim(m.melt) {
		return buf[:0], ErrAlreadyPlaintext
	}

	length := Rakaly_melt_data_length(m.melt)
	buf = resize(buf, int(length)) //#nosec G115
	if Rakaly_melt_write_data(m.melt, unsafe.SliceData(buf), length) != length {
		return buf[:0], ErrCopyFailed
	}

	return buf, nil
}

func (m MeltedOutput) HasUnknownTokens() bool {
	return Rakaly_melt_binary_unknown_tokens(m.melt)
}

func (m MeltedOutput) Free() {
	Rakaly_free_melt(m.melt)
}

type GameFile struct {
	file PdsFile
}

func (g GameFile) IsBinary() bool {
	return Rakaly_file_is_binary(g.file)
}

func (g GameFile) MeltMeta() (MeltedOutput, error) {
	meta := Rakaly_file_meta(g.file)
	if meta == nil {
		return MeltedOutput{}, ErrMetadataNotFound
	}

	meltResult := Rakaly_file_meta_melt(meta)
	err := unwrapError(Rakaly_melt_error(meltResult))
	if err != nil {
		return MeltedOutput{}, err
	}

	melt := Rakaly_melt_value(meltResult)
	return MeltedOutput{melt}, nil
}

func (g GameFile) Melt() (MeltedOutput, error) {
	meltResult := Rakaly_file_melt(g.file)
	err := unwrapError(Rakaly_melt_error(meltResult))
	if err != nil {
		return MeltedOutput{}, err
	}

	melt := Rakaly_melt_value(meltResult)
	return MeltedOutput{melt}, nil
}

func (g GameFile) Free() {
	Rakaly_free_file(g.file)
}

func ParseEu4(data []byte) (GameFile, error) {
	fileResult := Rakaly_eu4_file(unsafe.SliceData(data), uint(len(data)))
	err := unwrapError(Rakaly_file_error(fileResult))
	if err != nil {
		return GameFile{}, err
	}

	file := Rakaly_file_value(fileResult)
	return GameFile{file}, nil
}

func ParseCk3(data []byte) (GameFile, error) {
	fileResult := Rakaly_ck3_file(unsafe.SliceData(data), uint(len(data)))
	err := unwrapError(Rakaly_file_error(fileResult))
	if err != nil {
		return GameFile{}, err
	}

	file := Rakaly_file_value(fileResult)
	return GameFile{file}, nil
}

func ParseImperator(data []byte) (GameFile, error) {
	fileResult := Rakaly_imperator_file(unsafe.SliceData(data), uint(len(data)))
	err := unwrapError(Rakaly_file_error(fileResult))
	if err != nil {
		return GameFile{}, err
	}

	file := Rakaly_file_value(fileResult)
	return GameFile{file}, nil
}

func ParseHoi4(data []byte) (GameFile, error) {
	fileResult := Rakaly_hoi4_file(unsafe.SliceData(data), uint(len(data)))
	err := unwrapError(Rakaly_file_error(fileResult))
	if err != nil {
		return GameFile{}, err
	}

	file := Rakaly_file_value(fileResult)
	return GameFile{file}, nil
}

func ParseVic3(data []byte) (GameFile, error) {
	fileResult := Rakaly_vic3_file(unsafe.SliceData(data), uint(len(data)))
	err := unwrapError(Rakaly_file_error(fileResult))
	if err != nil {
		return GameFile{}, err
	}

	file := Rakaly_file_value(fileResult)
	return GameFile{file}, nil
}
