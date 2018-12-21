package main

import "fmt"

func main() {
	var primeTable [100]int
	primeTable[0] = 2 //素数を格納する配列。0番目の要素を２に初期化
	primeSize := 1    //格納されている素数の個数

	//奇数列を生成
	for n := 3; n <= len(primeTable); n += 2 {
		isPrime := true //ポイント
		for i := 1; i < primeSize; i++ {
			p := primeTable[i]
			if p*p > n {
				break
			}
			if n%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeTable[primeSize] = n
			primeSize++
		}
	}
	for i := 0; i < primeSize; i++ {
		fmt.Print(primeTable[i], " ")
	}
	fmt.Print("\n")

}
