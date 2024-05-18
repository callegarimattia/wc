//go:build unit

package wc_test

import (
	_ "embed"
	"fmt"
	"os/exec"
	"testing"

	"ccwc/internal/wc"
)

//go:embed testdata/test.txt
var testFile []byte

type Expected struct {
	Lines int
	Words int
	Bytes int
}

func TestCount(t *testing.T) {
	t.Parallel()
	tester := wc.New()

	cmd := exec.Command("wc", "testdata/test.txt")
	stdout, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	want := Expected{}
	_, err = fmt.Sscanf(string(stdout), "%d %d %d", &want.Lines, &want.Words, &want.Bytes)

	t.Run("check bytes", func(t *testing.T) {
		t.Parallel()
		result := tester.Count(testFile)
		if result.Bytes != want.Bytes {
			t.Errorf("Expected %v bytes, got %d", want.Bytes, result.Bytes)
		}
	})

	t.Run("check lines", func(t *testing.T) {
		t.Parallel()
		result := tester.Count(testFile)
		if result.Lines != want.Lines {
			t.Errorf("Expected %v lines, got %d", want.Lines, result.Lines)
		}
	})

	t.Run("check words", func(t *testing.T) {
		t.Parallel()
		result := tester.Count(testFile)
		if result.Words != want.Words {
			t.Errorf("Expected %v words, got %d", want.Words, result.Words)
		}
	})
}
