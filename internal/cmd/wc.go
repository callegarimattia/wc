package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"ccwc/internal/wc"
)

type Flags struct {
	line bool
	word bool
	byte bool
}

const (
	lineFlag = "l"
	wordFlag = "w"
	byteFlag = "c"
)

func WC(args []string, w io.Writer) {
	flagsRaw := readFlags(args)
	flags := parseFlags(flagsRaw)
	filesPath := readFilesPath(args)

	counter := wc.New()

	results := []wc.Result{}
	for _, path := range filesPath {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		results = append(results, counter.Count(data))
	}

	for i, res := range results {
		printOutput(w, filesPath[i], res, flags)
	}

	if len(results) > 1 {
		total := wc.Result{}
		for _, res := range results {
			total.Lines += res.Lines
			total.Words += res.Words
			total.Bytes += res.Bytes
		}
		printOutput(w, "total", total, flags)
	}

	if len(filesPath) == 0 {
		data := readFromStdin()
		res := counter.Count(data)
		printOutput(w, "", res, flags)
	}
}

func readFromStdin() []byte {
	data := []byte{}
	for {
		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}
		data = append(data, buf[:n]...)
	}
	return data
}

func printOutput(w io.Writer, path string, res wc.Result, flags Flags) {
	output := strings.Builder{}
	output.WriteString("  ")
	if flags.line {
		output.WriteString("  ")
		output.WriteString(strconv.Itoa(res.Lines))
	}
	if flags.word {
		output.WriteString("  ")
		output.WriteString(strconv.Itoa(res.Words))
	}
	if flags.byte {
		output.WriteString("  ")
		output.WriteString(strconv.Itoa(res.Bytes))
	}
	output.WriteString(" ")
	output.WriteString(path)
	fmt.Fprintln(w, output.String())
}

func readFilesPath(args []string) []string {
	filesPath := []string{}
	for len(args) > 0 {
		filesPath = append(filesPath, args[0])
		args = args[1:]
	}
	return filesPath
}

func readFlags(args []string) []string {
	flags := []string{}
	for len(args) > 0 {
		if args[0][0] != '-' {
			break
		}
		for _, flag := range args[0][1:] {
			flags = append(flags, string(flag))
		}
		args = args[1:]
	}
	return flags
}

func parseFlags(input []string) Flags {
	flags := Flags{}
	if len(input) == 0 {
		return Flags{true, true, true}
	}
	for _, flag := range input {
		switch flag {
		case lineFlag:
			flags.line = true
		case wordFlag:
			flags.word = true
		case byteFlag:
			flags.byte = true
		}
	}
	return flags
}
