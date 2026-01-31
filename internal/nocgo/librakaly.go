// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package nocgo

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

type (
	MeltedBuffer       *struct{}
	MeltedBufferResult *struct{}
	PdsError           *struct{}
	PdsFile            *struct{}
	PdsFileResult      *struct{}
	PdsMeta            *struct{}
)

var (
	Rakaly_free_melt                  func(MeltedBuffer)
	Rakaly_melt_data_length           func(MeltedBuffer) uint
	Rakaly_melt_is_verbatim           func(MeltedBuffer) bool
	Rakaly_melt_binary_unknown_tokens func(MeltedBuffer) bool
	Rakaly_melt_write_data            func(MeltedBuffer, *byte, uint) uint
	Rakaly_file_error                 func(PdsFileResult) PdsError
	Rakaly_error_length               func(PdsError) int
	Rakaly_error_write_data           func(PdsError, *byte, int) int
	Rakaly_free_error                 func(PdsError)
	Rakaly_free_file                  func(PdsFile)
	Rakaly_file_value                 func(PdsFileResult) PdsFile
	Rakaly_file_meta                  func(PdsFile) PdsMeta
	Rakaly_file_meta_melt             func(PdsMeta) MeltedBufferResult
	Rakaly_file_melt                  func(PdsFile) MeltedBufferResult
	Rakaly_file_is_binary             func(PdsFile) bool
	Rakaly_melt_error                 func(MeltedBufferResult) PdsError
	Rakaly_melt_value                 func(MeltedBufferResult) MeltedBuffer
	Rakaly_eu4_file                   func(*byte, uint) PdsFileResult
	Rakaly_ck3_file                   func(*byte, uint) PdsFileResult
	Rakaly_imperator_file             func(*byte, uint) PdsFileResult
	Rakaly_hoi4_file                  func(*byte, uint) PdsFileResult
	Rakaly_vic3_file                  func(*byte, uint) PdsFileResult
	Rakaly_eu5_file                   func(*byte, uint) PdsFileResult
)

func init() {
	var lib uintptr
	var err error
	for _, libDir := range filepath.SplitList(os.Getenv("LD_LIBRARY_PATH")) {
		libDir, err = filepath.EvalSymlinks(libDir)
		if err != nil {
			continue
		}
		libDir, err = filepath.Abs(libDir)
		if err != nil {
			continue
		}
		lib, err = loadLibrakaly(libDir)
		if err == nil {
			break
		}
	}
	if lib == 0 {
		executablePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		executablePath, err = filepath.EvalSymlinks(executablePath)
		if err != nil {
			panic(err)
		}
		libDir := filepath.Dir(executablePath)
		lib, err = loadLibrakaly(libDir)
		if err != nil {
			panic(err)
		}
	}

	purego.RegisterLibFunc(&Rakaly_free_melt, lib, "rakaly_free_melt")
	purego.RegisterLibFunc(&Rakaly_melt_data_length, lib, "rakaly_melt_data_length")
	purego.RegisterLibFunc(&Rakaly_melt_is_verbatim, lib, "rakaly_melt_is_verbatim")
	purego.RegisterLibFunc(&Rakaly_melt_binary_unknown_tokens, lib, "rakaly_melt_binary_unknown_tokens")
	purego.RegisterLibFunc(&Rakaly_melt_write_data, lib, "rakaly_melt_write_data")
	purego.RegisterLibFunc(&Rakaly_file_error, lib, "rakaly_file_error")
	purego.RegisterLibFunc(&Rakaly_error_length, lib, "rakaly_error_length")
	purego.RegisterLibFunc(&Rakaly_error_write_data, lib, "rakaly_error_write_data")
	purego.RegisterLibFunc(&Rakaly_free_error, lib, "rakaly_free_error")
	purego.RegisterLibFunc(&Rakaly_free_file, lib, "rakaly_free_file")
	purego.RegisterLibFunc(&Rakaly_file_value, lib, "rakaly_file_value")
	purego.RegisterLibFunc(&Rakaly_file_meta, lib, "rakaly_file_meta")
	purego.RegisterLibFunc(&Rakaly_file_meta_melt, lib, "rakaly_file_meta_melt")
	purego.RegisterLibFunc(&Rakaly_file_melt, lib, "rakaly_file_melt")
	purego.RegisterLibFunc(&Rakaly_file_is_binary, lib, "rakaly_file_is_binary")
	purego.RegisterLibFunc(&Rakaly_melt_error, lib, "rakaly_melt_error")
	purego.RegisterLibFunc(&Rakaly_melt_value, lib, "rakaly_melt_value")
	purego.RegisterLibFunc(&Rakaly_eu4_file, lib, "rakaly_eu4_file")
	purego.RegisterLibFunc(&Rakaly_ck3_file, lib, "rakaly_ck3_file")
	purego.RegisterLibFunc(&Rakaly_imperator_file, lib, "rakaly_imperator_file")
	purego.RegisterLibFunc(&Rakaly_hoi4_file, lib, "rakaly_hoi4_file")
	purego.RegisterLibFunc(&Rakaly_vic3_file, lib, "rakaly_vic3_file")
	purego.RegisterLibFunc(&Rakaly_eu5_file, lib, "rakaly_eu5_file")
}

var libName = func() string {
	switch runtime.GOOS {
	case "darwin":
		return "librakaly.dylib"
	case "linux":
		return "librakaly.so"
	case "windows":
		return "rakaly.dll"
	default:
		panic("GOOS=" + runtime.GOOS + " is not supported")
	}
}()

func loadLibrakaly(libDir string) (uintptr, error) {
	libPath := libDir + string(filepath.Separator) + libName
	lib, err := openLibrary(libPath)
	if err != nil {
		return 0, errors.New("failed to load library (" + libName + "): " + err.Error())
	}
	return lib, nil
}
