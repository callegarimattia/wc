package wc

// #include <wctype.h>
// #include <wchar.h>
import "C"

type wc struct{}

type Counter interface {
	Count([]byte) Result
}

type Result struct {
	Lines int
	Words int
	Bytes int
}

func New() Counter {
	return &wc{}
}

func (w *wc) Count(data []byte) Result {
	bytes, words, lines := 0, 0, 0

	var inWord bool
	for i := 0; i < len(data); i++ {
		bytes++
		if data[i] == '\n' {
			lines++
		}
		if C.iswspace(C.wint_t(data[i])) != 0 {
			inWord = false
		} else if !inWord {
			inWord = true
			words++
		}
	}

	return Result{lines, words, bytes}
}
