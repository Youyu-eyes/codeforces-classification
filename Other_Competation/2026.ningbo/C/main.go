package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func II() int {
	n := 0
	neg := false
	b, _ := in.ReadByte()
	for b < '0' || b > '9' {
		if b == '-' {
			neg = true
			b, _ = in.ReadByte()
			break
		}
		b, _ = in.ReadByte()
	}
	for b >= '0' && b <= '9' {
		n = n*10 + int(b&15)
		b, _ = in.ReadByte()
	}
	if neg {
		n = -n
	}
	return n
}

func Println(a ...any) {
	fmt.Fprintln(out, a...)
	out.Flush()
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	T := II()
	for t := 0; t < T; t++ {
		solve()
	}
	out.Flush()
}
