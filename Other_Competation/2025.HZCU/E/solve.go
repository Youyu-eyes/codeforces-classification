// 非暴力做法
// 字符串哈希
// 复杂度：O(n)

// 原题 O(nk) 可过，k 为目标串长度，本题为 3
// 仍有 KMP 做法

package main

func solve() {
	S := []byte(IS())
	T := []byte("hyw")
	ns, nt := len(S), len(T)
	s := NewDoubleHash(S)
	t := NewDoubleHash(T)
	tVal := t.GetHash(0, nt - 1)

	ans := 0

	for i := 0; i <= ns - nt; i++ {
		if s.GetHash(i, i + nt - 1) == tVal {
			ans ^= (i + 1)
		}
	}

	Println(ans)
}

type IntegerHash interface {
	~int | ~int64 | ~byte | ~rune
}

const (
	Base1 = 1_000_000_021
	Mod1  = 1_000_000_007
	Base2 = 1_000_000_033
	Mod2  = 1_000_000_009
)

type DoubleHash struct {
	pref1, pref2 []int
	pow1, pow2   []int
}

type HashVal struct {
	h1 int
	h2 int
}

// 字符串双哈希模板
// 支持 int，byte，rune 类型
// O(1) 时间返回任意子数组的双哈希值 HashVal

// 可以用于判断两个不同字符串的子串是否相等

func NewDoubleHash[T IntegerHash](arr []T) *DoubleHash {
	n := len(arr)
	h := &DoubleHash{
		pref1: make([]int, n+1),
		pref2: make([]int, n+1),
		pow1:  make([]int, n+1),
		pow2:  make([]int, n+1),
	}
	h.pow1[0], h.pow2[0] = 1, 1

	for i, val := range arr {
		v := int(val)

		h.pow1[i+1] = h.pow1[i] * Base1 % Mod1
		h.pow2[i+1] = h.pow2[i] * Base2 % Mod2

		h.pref1[i+1] = (h.pref1[i]*Base1 + v) % Mod1
		h.pref2[i+1] = (h.pref2[i]*Base2 + v) % Mod2
	}
	return h
}

func (h *DoubleHash) GetHash(l, r int) HashVal {
	if l > r || l < 0 || r >= len(h.pref1)-1 {
		return HashVal{-1, -1}
	}
	length := r - l + 1
	
	hash1 := (h.pref1[r+1] - h.pref1[l]*h.pow1[length]) % Mod1
	if hash1 < 0 {
		hash1 += Mod1
	}

	hash2 := (h.pref2[r+1] - h.pref2[l]*h.pow2[length]) % Mod2
	if hash2 < 0 {
		hash2 += Mod2
	}
	
	return HashVal{hash1, hash2}
}

// CompareSubstrings 比较 arr[l1: r1+1] 和 arr[l2: r2+1] 的字典序
// 返回 -1 表示 子串1 < 子串2
// 返回  0 表示 子串1 == 子串2
// 返回  1 表示 子串1 > 子串2
func CompareSubstrings[T IntegerHash](h *DoubleHash, arr []T, l1, r1, l2, r2 int) int {
	len1 := r1 - l1 + 1
	len2 := r2 - l2 + 1
	
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	// 双开区间二分 (l, r)
	left := 0          // l 维护“匹配成功”的边界，初始 0 必定成功
	right := minLen + 1 // r 维护“匹配失败”的边界，初始 minLen+1 必定失败
	
	for left + 1 < right {
		mid := left + (right - left) >> 1
		
		// 检查长度为 mid 的前缀哈希是否相等
		if h.GetHash(l1, l1 + mid - 1) == h.GetHash(l2, l2 + mid - 1) {
			left = mid
		} else {
			right = mid
		}
	}
	
	lcp := left

	if lcp == len1 && lcp == len2 {
		return 0
	}

	if lcp == len1 {
		return -1
	}

	if lcp == len2 {
		return 1
	}

	if arr[l1 + lcp] < arr[l2 + lcp] {
		return -1
	}
	return 1
}
