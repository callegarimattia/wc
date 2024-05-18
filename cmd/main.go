package main

import (
	"os"

	"ccwc/internal/cmd"
)

func main() {
	cmd.WC(os.Args[1:], os.Stdout)
}
