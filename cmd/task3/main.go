package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Index       int
	Letter      rune
	ReturnCount int
}

type Queue struct {
	queue []*Node
	maxN  int
	head  int
	tail  int
	size  int
}

func NewQueue(n int) *Queue {
	return &Queue{
		queue: make([]*Node, n),
		maxN:  n,
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(x *Node) {
	if q.size != q.maxN {
		q.queue[q.tail] = x
		q.tail = (q.tail + 1) % q.maxN
		q.size += 1
	}
}

func (q *Queue) Pop() *Node {
	if q.IsEmpty() {
		return nil
	}
	x := q.queue[q.head]
	q.queue[q.head] = nil
	q.head = (q.head + 1) % q.maxN
	q.size -= 1
	return x
}

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

		queue := NewQueue(len(runeArray))

		testWord := strings.Repeat(string(runeArray[0]), len(runeArray))
		runeTestWord := []rune(testWord)
		for j := 1; j < len(word)-1; j++ {
			if runeArray[j] != runeTestWord[0] {
				queue.Push(&Node{
					Index:  j,
					Letter: runeArray[j],
				})
			}
		}

		if newWord, ok := tryToSetLtr(runeTestWord, *queue); !ok {
			_, err := out.WriteString("NO" + "\n")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			continue
		} else {
			if string(runeArray) == string(newWord) {
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

	}
}

func tryToSetLtr(word []rune, queue Queue) ([]rune, bool) {
	for {
		if queue.IsEmpty() {
			return word, true
		}
		node := queue.Pop()
		if node.ReturnCount > 2 {
			return word, false
		}

		if word[node.Index+1] == word[node.Index-1] {
			word[node.Index] = node.Letter
		} else {
			node.ReturnCount++
			queue.Push(node)
		}
	}
}
