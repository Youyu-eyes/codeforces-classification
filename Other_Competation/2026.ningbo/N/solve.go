package main

// ------- 全局预处理变量 ------- //
var (
	isPrimeArr []bool
	primes     []int
)

// func init() 会在 main 函数之前自动执行，用于预处理各类数据
func init() {
	isPrimeArr = make([]bool, MX)

	for i := range isPrimeArr {
		isPrimeArr[i] = true
	}
	isPrimeArr[0], isPrimeArr[1] = false, false
	for i := 2; i < MX; i++ {
		if isPrimeArr[i] {
			primes = append(primes, i)
			for j := i * i; j < MX; j += i {
				isPrimeArr[j] = false // j 是质数 i 的倍数
			}
		}
	}
}

func solve() {
	n := II()
	Print(2)
	Print(" ")
	Print(primes[n - 1])
}