package main

import (
	"flag"
	"fmt"
)

func findSubstring(str, substr string) bool {
	strRunes := []rune(str)
	substrRunes := []rune(substr)

	for i := range strRunes[:len(strRunes)-len(substrRunes)+1] {
		match := true

		for j := range substrRunes {

			if strRunes[i+j] != substrRunes[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println()
	fmt.Println("Программа для нахождения подстроки в кириллической строке.")
	fmt.Println()
	strPtr := flag.String("str", "Программирование это просто", "the string to search in")
	substrPtr := flag.String("substr", "просто", "the substring to search for")
	flag.Parse()

	if *strPtr == "" || *substrPtr == "" {
		fmt.Println("Both -str and -substr flags are required")
		return
	}

	if findSubstring(*strPtr, *substrPtr) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
