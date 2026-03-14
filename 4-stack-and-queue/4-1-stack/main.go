package main

import "fmt"

// 数値を扱うスタック構造体
type IntStack struct {
	stk      []int // スタック配列本体となるスライス
	capacity int   // スタックの容量
	ptr      int   // スタックポインタ、スタックに積まれているデータの個数を表す
}

func NewIntStack(maxlen int) *IntStack {
	return &IntStack{
		capacity: maxlen,
		stk:      make([]int, maxlen),
		ptr:      0,
	}
}

func (s *IntStack) Push(x int) error {
	if s.ptr >= s.capacity {
		return fmt.Errorf("スタックが満杯です")
	}

	s.stk[s.ptr] = x
	s.ptr++
	return nil
}

func (s *IntStack) Pop() (int, error) {
	if s.ptr <= 0 {
		return 0, fmt.Errorf("スタックが空です")
	}
	s.ptr--
	return s.stk[s.ptr], nil
}

func main() {
	stack := NewIntStack(5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	for _, s := range stack.stk {
		fmt.Printf("ptr: %d, value: %d\n", stack.ptr, s)
	}

	stack.Pop()
	for _, s := range stack.stk {
		fmt.Printf("ptr: %d, value: %d\n", stack.ptr, s)
	}
}
