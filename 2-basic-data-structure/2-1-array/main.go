package main

import (
	"bufio"
	"fmt"
	"os"
)

var sreader = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println("身長の最大値を求めます")
	fmt.Println("人数は：")
	var count int
	fmt.Fscan(sreader, &count)

	heights := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("%d人目の身長は：", i+1)
		fmt.Fscan(sreader, &heights[i])
	}

	fmt.Printf("身長の最大値は%dです\n", maxOfArray(heights))
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
