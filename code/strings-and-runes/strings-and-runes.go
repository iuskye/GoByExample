package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s 是一个 string 分配了一个 literal value 表示泰语中的单词 “hello”。Go 字符串是
	// UTF-8编码的文本。
	const s = "สวัสดี"

	// 因为字符串等价于 []byte，这会产生存储在其中的原始字节长度。
	fmt.Println("Len:", len(s))

	// 对字符串进行索引会在每个索引处生成原始字节值。这个循环生成构成 s 中 Unicode 的
	// 所有字节的十六进制值。
	for i := 0; i < len(s); i++ {
		fmt.Println("%x:", s[i])
	}
	fmt.Println()

	// 要计算字符串有多少 rune，我们可以使用 utf8 包。注意 RuneCountInString 的运行时
	// 取决于字符串的大笑。因为它必须按顺序解码每个 UTF-8 rune。一些泰语字符由多个 UTF-8
	// code point 表示，所以这个计数的结果可能会令人惊讶。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串并解码每个 rune 及其在字符串中的偏移量。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 我们可以通过显式使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// 这演示了讲 rune value 传递给函数。
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// 用单引号括起来的值是 rune literals。我们可以直接讲 rune value 与 rune
	// literal 进行比较。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
