package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanWords)

	in.Scan()
	t, _ := strconv.Atoi(in.Text())

	for i := 0; i < t; i++ {
		in.Scan()
		word := in.Text()

		runeArray := []rune(word)

		if len(runeArray) == 1 {
			_, err := out.WriteString("YES" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			continue
		}

		if len(runeArray) == 2 {
			if runeArray[0] == runeArray[1] {
				_, err := out.WriteString("YES" + "\n")
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				continue
			} else {
				_, err := out.WriteString("NO" + "\n")
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				continue
			}
		}

		if runeArray[0] != runeArray[len(runeArray)-1] {
			_, err := out.WriteString("NO" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			continue
		}

		if test(runeArray) {
			_, err := out.WriteString("YES" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			_, err := out.WriteString("NO" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
}

func test(runeArray []rune) bool {
	for j := 1; j < len(runeArray)-2; j++ {
		if runeArray[j-1] != runeArray[j+1] && runeArray[j] != runeArray[0] {
			return false
		}
	}

	return true
}
