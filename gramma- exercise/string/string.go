package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s))

	// 普通 for循环取得是ascii码
	for i := 0; i < len(s); i++ {
		// 16进制
		fmt.Printf("Ox%x ", s[i])
	}
	fmt.Printf("\n")

	// RuneCountInString()用来统计Unicode字符数量
	fmt.Println("the character count in s is", utf8.RuneCountInString(s))
	for _, c := range s {
		fmt.Printf("Ox%x ", c) // Ox4e2d Ox56fd Ox4eba
	}
	fmt.Printf("\n")

	encodeRune()

}

func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode charactor is %c\n", r)
	buf := make([]byte, 3)

	// 对 rune 进行utf-8编码
	_ = utf8.EncodeRune(buf, r)
	fmt.Printf("utf-8 representation is 0x%X\n", buf)
}
