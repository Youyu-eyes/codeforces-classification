package main

import (
	"strconv"
)

func solve() {
	Println("? 1")
	res := II()
	if res == 0 {
		Println("! 2")
		return
	}

	cur := 1
	lastGood := 1

	for {
		if cur > 1_000_000_000_000_000_000 {
			Println("! " + strconv.Itoa(lastGood))
			return
		}

		Println("? " + strconv.Itoa(cur))
		res = II()

		if res == 0 {
			bad := cur
			good := cur / 2
			l := good + 1
			r := bad - 1
			ans := good

			for l <= r {
				m := (l + r) / 2
				Println("? " + strconv.Itoa(m))
				res2 := II()
				if res2 == 1 {
					ans = m
					l = m + 1
				} else {
					r = m - 1
				}
			}

			if ans == good {
				if bad+1 <= 1_000_000_000_000_000_000 {
					Println("? " + strconv.Itoa(bad+1))
					if II() == 1 {
						Println("! " + strconv.Itoa(bad+1))
						return
					}
				}
				if bad+2 <= 1_000_000_000_000_000_000 {
					Println("? " + strconv.Itoa(bad+2))
					if II() == 1 {
						Println("! " + strconv.Itoa(bad+2))
						return
					}
				}
				Println("! " + strconv.Itoa(good))
				return
			} else {
				n1 := 2*ans + 1
				n2 := 2*ans + 2

				if n1 <= 1_000_000_000_000_000_000 {
					Println("? " + strconv.Itoa(n1))
					if II() == 1 {
						Println("! " + strconv.Itoa(n1))
						return
					}
				}
				Println("! " + strconv.Itoa(n2))
				return
			}
		}

		lastGood = cur
		if cur > 500_000_000_000_000_000 {
			Println("! " + strconv.Itoa(cur))
			return
		}
		cur *= 2
	}
}