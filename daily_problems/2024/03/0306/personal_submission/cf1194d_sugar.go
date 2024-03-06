package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func cf1194d(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n, k int
		Fscan(in, &n, &k)
		if k%3 != 0 {
			if n%3 == 0 {
				Fprintln(out, "Bob")
			} else {
				Fprintln(out, "Alice")
			}
			continue
		}
		mod := k + 1
		n %= mod
		if n == mod-1 || n%3 != 0 {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

func main() {
	cf1194d(os.Stdin, os.Stdout)
}
