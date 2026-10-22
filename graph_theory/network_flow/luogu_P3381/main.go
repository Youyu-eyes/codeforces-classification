package main

import (
	"bufio"
	"fmt"
	"os"	
)

const (
	MOD = 1_000_000_007
	inf = 1 << 60
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

// n := II()
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

// n := IB()
func IB() byte {
    for {
        b, err := in.ReadByte()
        if err != nil {
            return 0 // 或 panic，按需求处理
        }
        if b != ' ' && b != '\n' && b != '\t' && b != '\r' {
            return b
        }
    }
}

// n := IF()
func IF() (x float64) {
	fmt.Fscan(in, &x)
	return
}

// s := IS()
func IS() (x string) {
	fmt.Fscan(in, &x)
	return
}

// 读入带空格的字符串 
// s, _ := in.ReadString('\n')
// s = strings.TrimRight(s, "\n\r")

func Print(a ...any)            { fmt.Fprint(out, a...) }
func Println(a ...any)          { fmt.Fprintln(out, a...) }
func Printf(f string, a ...any) { fmt.Fprintf(out, f, a...) }

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := 1

	for t := 0; t < T; t++ {
		solve()
	}
}

func abs[T int | int64 | float64](x T) T {
	if x >= 0 { return x }
	return -x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
