package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func f(v interface{}) {
	fmt.Printf("%#v\n", v)
}
func main() {
	// 字符串
	aString := "Hello World! 王 健 €"
	fmt.Println("First character", string(aString[0]))
	r := '€'
	fmt.Println("As an int32 value:", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	for _, v := range aString {
		fmt.Printf("%c", v) // 如果使用rune类型, 使用%c
	}
	fmt.Println()
	// 字符串和数字类型转化
	// 1. strconv.Itoa和strconv.FormatInt: 按照相同的表现和项目的字符数转化
	// 2. string(): 转化为Unicode code point, 是一个单字符, 例如100是ASCII码d
	n := 100
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa: %#v\n", input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.Itoa: %#v\n", input)
	input = string(n)
	fmt.Printf("strconv.Itoa: %#v\n", input)
	// strings核心包
	fmt.Printf("%#v\n", strings.Fields("This is  a string\t!")) // unicode.IsSpace(): '\t', '\n', '\v', '\f', '\r', ' '
	f(strings.Split("abcd efg", ""))
	f(strings.Replace("abcd efg", "", "_", -1))
	f(strings.Split("123++432++", "++"))
	f(strings.SplitAfter("123++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f(strings.TrimFunc("123 abc 23 ABC \t .", trimFunction))
	// 时间和日期
	// 1. 时间精度纳秒
	f(time.Now().Unix())
	f(time.Unix(time.Now().Unix(), 0))
	f(time.Now().Format("2006-01-02 15:04:05")) // 年(6)月(1)日(2)时(3/15)分(4)秒(5)
	current_time := time.Now()
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())
	d, err := time.Parse("2006-01-02 15:04:05", "2022-01-07 09:37:01")
	if err != nil {
		fmt.Print(err)
		return
	}
	f(d)
	d = time.Unix(time.Now().Unix(), 0)
	duration := time.Since(current_time)
	fmt.Println("Execution time:", duration)
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", current_time.In(loc))
}
