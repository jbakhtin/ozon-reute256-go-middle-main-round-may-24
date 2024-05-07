package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int

	_, err := fmt.Fscan(in, &a, &b)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Fprintln(out, a-b)
	if err != nil {
		os.Exit(1)
	}
}
