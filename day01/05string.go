package main

import (
	"fmt"
	"strings"
)

// 用双引号包裹的是字符串
// 用单引号包裹的是字符
// 用``包裹的字符串为多行字符串

func main() {
	s1 := "aaa-bbb"
	s2 := 'a'
	s3 := '你'
	s4 := "你"

	fmt.Println(s1, s2, s3, s4)

	ret := strings.Split(s1, "-")
	fmt.Println(ret)

	fmt.Println(strings.Contains(s1, "-"))
	fmt.Println(strings.HasPrefix(s1, "a"))
	fmt.Println(strings.HasSuffix(s1, "bb"))
	fmt.Println(strings.Index(s1, "a"))
	fmt.Println(strings.Join(ret, "="))

	// 字符串修改
	s5 := "红绿黄"
	// 只是简单字符可使用byte
	r1 := []rune(s5)
	r1[0] = '黑'
	fmt.Println(string(r1))
}