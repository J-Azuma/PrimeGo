package main

import "fmt"

func isPrime(n, primeSize int, primeTable []int) bool {
	for i := 1; i < primeSize; i++ {
		p := primeTable[i]
		if p*p > n {
			break
		}
		if n%p == 0 {
			return false
		}
	}
	return true
}

func main() {
	primeTable := make([]int, 100)
	primeTable[0] = 2
	primeSize := 1 //18 19行目で配列及び配列の個数を初期化
	for n := 3; n < len(primeTable); n += 2 {
		if isPrime(n, primeSize, primeTable) {
			primeTable[primeSize] = n
			primeSize++
		}
	}
	fmt.Println(primeTable[:primeSize]) //primeSize個の素数からなる配列primeTableを出力
}
