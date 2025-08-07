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

type TestCase struct {
	URL       string
	ParseFunc ParseFunc
}

type ParseFunc func(data []byte) (rakaly.GameFile, error)

func (t *TestCase) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(t.ParseFunc).Pointer()).Name()
	name = strings.TrimPrefix(name, "github.com/antoniszymanski/rakaly-go.Parse")
	return name
}

func TestMain(t *testing.T) {
	for _, testCase := range []TestCase{
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
		t.Run(testCase.Name(), func(t *testing.T) {
			RunTest(t, testCase)
		})
	}
}

func RunTest(t *testing.T, testCase TestCase) {
	t.Parallel()

	resp, err := http.Get(testCase.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if testCase.Name() == "Hoi4" {
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

	g, err := testCase.ParseFunc(data)
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
