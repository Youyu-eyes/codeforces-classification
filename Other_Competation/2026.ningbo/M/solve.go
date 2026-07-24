package main

func solve() {
	S := IS()
	s := []byte(S)
	st := []byte{}
	ans := 0

	del := func() bool {
		if L := len(st); L >= 3 && st[L - 3] == 'I' && st[L - 2] == 'C' && st[L - 1] == 'P' {
			st = st[:L - 3]
			ans++
			return true
		}
		return false
	}

	for _, char := range s {
		if char == 'I' {
			st = append(st, 'I')
		} else if char == 'C' {
			if !del() {
				st = append(st, 'C')
			}
		} else {
			cnt := 0
			if len(st) == 0 || st[len(st) - 1] != 'I' {
				st = append(st, 'P')
				continue
			}
			for len(st) > 0 && st[len(st) - 1] == 'I' {
				cnt++
				st = st[:len(st) - 1]
			}
			st = append(st, 'P')
			for cnt > 0 && del() {
				cnt--
			}
			for i := 0; i < cnt; i++ {
                st = append(st, 'C')
            }
		}
	}

	Println(ans)
}