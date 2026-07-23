package main

import (
	"strings"
)

func solve() {
    n := II()

    v := make([]string, n)
    for i := 0; i < n; i++ {
        v[i] = IS()
    }

	in.ReadString('\n')
    s, _ := in.ReadString('\n')
    s = strings.TrimRight(s, "\n\r")

    words := strings.Split(s, " ")
    set := make(map[string]struct{}, len(words))
    for _, w := range words {
        set[w] = struct{}{}
    }

    for _, verb := range v {
        _, hasVerb := set[verb]
        _, hasVerbED := set[verb+"ed"]
        if hasVerb && hasVerbED {
            Println("not friendly")
            return
        }
    }
    Println("friendly")
}