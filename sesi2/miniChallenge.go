package main

import "fmt"

func miniChallenge() {
	sentence := "selamat malam"
	charMap := map[string]int{}

	for _, char := range sentence {
		fmt.Printf("%c\n",char)
		charMap[string(char)]++
	}

	fmt.Println(charMap)
}