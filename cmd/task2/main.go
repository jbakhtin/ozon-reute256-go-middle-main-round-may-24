package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in *bufio.Scanner
	var out *bufio.Writer
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanWords)

	var n, t int

	in.Scan()
	n, err := strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	in.Scan()
	t, err = strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	alphabet := make(map[string]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		alphabet[in.Text()]++
	}

	var word string
	for i := 0; i < t; i++ {
		in.Scan()
		word = in.Text()

		if isContain(word, alphabet) {
			_, err := out.WriteString("YES" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			continue
		}

		_, err := out.WriteString("NO" + "\n")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func isContain(word string, alphabet map[string]int) bool {
	wordAlphabet := make(map[string]int)
	for i := 0; i < len(word); i++ {
		if _, ok := alphabet[string(word[i])]; !ok {
			return false
		}
		wordAlphabet[string(word[i])]++
	}

	for key, val := range alphabet {
		if val != wordAlphabet[key] {
			return false
		}
	}

	return true
}
