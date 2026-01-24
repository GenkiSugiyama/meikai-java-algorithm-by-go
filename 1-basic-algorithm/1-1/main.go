package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Println("スペース区切りで数値を入力してください ↓")
	fmt.Print("a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	max4 := max4(a, b, c, d)
	min3 := min3(a, b, c)
	min4 := min4(a, b, c, d)
	fmt.Printf("max4: %d, min3: %d, min4: %d\n", max4, min3, min4)

}

// 演習1-1
func max4(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

// 演習1-2
func min3(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// 演習1-3
func min4(a, b, c, d int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}
	return min
}
