package main

import "fmt"

// 常量
const pi = 3.1415926

// 和第一个相同
const (
	// d 会报错
	a = 100
	b
	c
)

const (
	// 从0开始
	d = iota
	e
	f
)

const (
	g = 10
	k
	h = iota
	i
)


const (
    m, n = iota + 1, iota + 2
	s, t = iota + 1, iota + 2
)


// 定义数量级

const (
	_ = iota
	KB = 1 << (10 * iota) // 2的10次方
	MB = 1 << (10 * iota) // 2的20次方
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	// 10 10 2 3
	fmt.Println(g, k, h, i)
	fmt.Println(m, n, s, t)
}