package main

import (
	"bufio"
	"fmt"
	"os"
)

var sreader = bufio.NewReader(os.Stdin)

func main() {

	// fmt.Println("身長の最大値を求めます")
	// fmt.Println("人数は：")
	// var count int
	// fmt.Fscan(sreader, &count)

	// heights := make([]int, count)

	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%d人目の身長は：", i+1)
	// 	fmt.Fscan(sreader, &heights[i])
	// }

	// fmt.Printf("身長の最大値は%dです\n", maxOfArray(heights))

	// fmt.Println("要素数は：")
	// var count int
	// fmt.Fscan(sreader, &count)

	// array := make([]int, count)

	// // arrayに乱数を入れる
	// r := rand.New(rand.NewSource(42))
	// for i := 0; i < count; i++ {
	// 	array[i] = r.Intn(10)
	// }

	// fmt.Printf("配列の要素は%vです\n", array)
	// fmt.Println("配列を逆順にします")
	// reverseArray(array)
	// fmt.Printf("反転後の配列の要素は%vです\n", array)

	// fmt.Println("配列の全要素の合計を求めます")
	// fmt.Printf("配列arrayの全要素の合計は%dです\n", sumOfArray(array))

	// fmt.Println("配列arrayをコピーして新しい配列newArrayを作成します")
	// newArray := copyArray(array)
	// fmt.Printf("新しい配列newArrayの要素は%vです\n", newArray)

	// fmt.Println("配列arrayの要素を新配列へ逆順にコピーします")
	// rArray := make([]int, len(array))
	// rCopyArray(array, rArray)
	// fmt.Printf("新しい配列rArrayの要素は%vです\n", rArray)

	// fmt.Println("10進数を基数変換します")
	// var x int
	// for {
	// 	fmt.Println("変換したい10進数xは：")
	// 	fmt.Fscan(sreader, &x)
	// 	if x < 0 {
	// 		fmt.Println("xは0以上の整数を入力してください")
	// 		continue
	// 	}
	// 	break
	// }

	// fmt.Println("基数rは：")
	// var f int
	// fmt.Fscan(sreader, &f)

	// d := make([]rune, 32)
	// digits := cardCnovert(x, f, d)
	// fmt.Printf("%dを基数%dの数に変換した数字文字の並びは%vで、桁数は%dです\n", x, f, string(d[:digits]), digits)

	findPrimeNumber()
	findPrimeNumber2()
}

func maxOfArray(array []int) int {
	max := array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	return max
}

// 演習2-2
func reverseArray(array []int) {
	for i := 0; i < len(array)/2; i++ {
		fmt.Printf("配列の要素は%vです\n", array)
		fmt.Printf("array[%d]とarray[%d]を入れ替えます\n", i, len(array)-1-i)
		t := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = t
	}
	fmt.Println("配列の反転が完了しました")
}

// 演習2-3
func sumOfArray(array []int) int {
	result := 0
	for _, value := range array {
		result += value
	}
	return result
}

// 演習2-4
func copyArray(array []int) []int {
	newArray := make([]int, len(array))
	for i, value := range array {
		newArray[i] = value
	}
	return newArray
}

// 演習2-5
func rCopyArray(array []int, newArray []int) {
	for i := len(array) - 1; i >= 0; i-- {
		newArray[len(array)-1-i] = array[i]
	}
}

// 演習2-6
// 10進数xをr進数に変換した数字文字の並びを配列dに格納して配列の桁数を返す関数
func cardCnovert(x, r int, d []rune) int {
	// 変換後の桁数
	digits := 0
	// 変換後の数を表す文字列
	dchar := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Printf("%d | %d \n", r, x)
	fmt.Println("  +-----")

	for x > 0 {
		// 10進数xを基数rで余った数をdcharから取り出して配列dに格納
		d[digits] = rune(dchar[x%r])
		// 上位桁の数を求めるためにxをrで割る
		x /= r
		digits++
		if x == 0 {
			fmt.Printf("    %d ... %c\n", x, d[digits-1])
			break
		}
		fmt.Printf("%d | %d ... %c\n", r, x, d[digits-1])
		fmt.Println("  +-----")
	}

	// 配列dの並び順で基数変換後の数を表すために配列dの要素を逆順にする
	for i := 0; i < digits/2; i++ {
		t := d[i]
		d[i] = d[digits-1-i]
		d[digits-1-i] = t
	}

	return digits
}

// 1000以下の素数を求めるために2から1000までの数を順番に割り算していく方法
func findPrimeNumber() {
	counter := 0
	// 1000以下の素数を求める
	for i := 2; i <= 1000; i++ {
		var j int
		// 2からi-1までの数でiが割り切れたらiは素数ではない
		for j = 2; j < i; j++ {
			counter++
			if i%j == 0 {
				break
			}
		}
		if j == i {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n割り算の回数は%d回です\n", counter)
}

// 1000以下の素数を求めるためのその数までの素数で割り算していく方法
func findPrimeNumber2() {
	counter := 0
	ptr := 0
	// 1000以下の素数を求めるための配列
	primes := make([]int, 500)
	primes[ptr] = 2
	ptr++

	// 3以上の奇数iの素数判定を行う
	for i := 3; i <= 1000; i += 2 {
		var j int
		for j = 1; j < ptr; j++ {
			counter++
			// 判定対象の奇数までの素数で割り切れたらiは素数ではない
			if i%primes[j] == 0 {
				break
			}
		}
		if j == ptr {
			primes[ptr] = i
			ptr++
		}
	}
	fmt.Printf("1000以下の素数は%vです\n", primes[:ptr])
	fmt.Printf("割り算の回数は%d回です\n", counter)
}

func findPrimeNumber3() {
	counter := 0
	ptr := 0
	primes := make([]int, 500)

	primes[ptr] = 2
	ptr++
	primes[ptr] = 3
	ptr++

	for n := 5; n <= 1000; n += 2 {
		flag := false
		for i := 1; primes[i]*primes[i] <= n; i++ {

		}
	}
}
