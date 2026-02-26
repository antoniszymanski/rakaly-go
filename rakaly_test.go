// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package rakaly_test

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/antoniszymanski/rakaly-go"
)

func Test(t *testing.T) {
	t.Parallel()
	for _, tc := range []TestCase{
		{
			URL:       "https://eu4saves-test-cases.s3.us-west-002.backblazeb2.com/kandy2.bin.eu4",
			ParseFunc: rakaly.ParseEu4,
		},
		{
			URL:       "https://ck3saves-test-cases.s3.us-west-002.backblazeb2.com/af_Munso_867_Ironman.ck3",
			ParseFunc: rakaly.ParseCk3,
		},
		{
			URL:       "https://imperator-test-cases.s3.us-west-002.backblazeb2.com/observer1.5.rome",
			ParseFunc: rakaly.ParseImperator,
		},
		{
			URL:       "https://hoi4saves-test-cases.s3.us-west-002.backblazeb2.com/1.10-ironman.zip",
			ParseFunc: rakaly.ParseHoi4,
		},
	} {
		t.Run(tc.Name(), tc.Run)
	}
}

type TestCase struct {
	URL       string
	ParseFunc func(data []byte) (rakaly.GameFile, error)
}

func (tc *TestCase) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(tc.ParseFunc).Pointer()).Name()
	name = strings.TrimPrefix(name, "github.com/antoniszymanski/rakaly-go.Parse")
	return name
}

func (tc *TestCase) Run(t *testing.T) {
	t.Parallel()

	resp, err := http.Get(tc.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if tc.Name() == "Hoi4" {
		r := bytes.NewReader(data)
		zr, err := zip.NewReader(r, r.Size())
		if err != nil {
			t.Fatal(err)
		}

		f, err := zr.Open("1.10-ironman.hoi4")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close() //nolint:errcheck

		data, err = io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
	}

	g, err := tc.ParseFunc(data)
	if err != nil {
		t.Fatal(err)
	}
	defer g.Free()

	t.Log("IsBinary:", g.IsBinary())

	m, err := g.Melt()
	if err != nil {
		t.Fatal(err)
	}
	defer m.Free()

	t.Log("HasUnknownTokens:", m.HasUnknownTokens())

	data, err = m.WriteData(data)
	if err != nil {
		t.Fatal(err)
	}
	_ = data
}
