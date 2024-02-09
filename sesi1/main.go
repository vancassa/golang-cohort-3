package main

import (
	"fmt"
	"strconv"
)

const (
	statusNew = iota +2
	statusOld
)

func main() {
	// Variable with data type
	var title string = "Golang cohort 3"
	var totalMember int = 20

	var title1 string
	title1 = "Golang Beginner"

	fmt.Println("Variable with data type:", title, totalMember, title1)

	// Variable without data type
	title2 := "Golang cohort 3"
	totalMember1 := 20
	fmt.Printf("Without data type: %T %T\n", title2, totalMember1)
	fmt.Printf("Without data type: %s %d\n", title2, totalMember1)

	// Multiple declaration 
	var title3, title4, title5 string = "satu", "dua", "tiga";
	var age1, age2, age3 int
	age1, age2, age3 = 1,2,3
	fmt.Println("Multiple declaration:", title3, title4, title5)
	fmt.Println("Multiple declaration:", age1,age2,age3)

	// Underscore variable
	var _ string = "Not used and program can be used without error"

	// Constants
	fmt.Println("statusNew, statusOld:", statusNew, statusOld)

	// ATOI
	stringNumber := "1234"
	num, _ := strconv.Atoi(stringNumber)
	fmt.Printf("stringNumber converted to type: %T \n", num)

	// Arithmetic
	var value = (2+2)*3
	fmt.Println("Arithmetic", value)

	// Relational
	var one bool = 2 < 3
	var two bool = "women" == "Women"
	var three bool = 10 != 2
	fmt.Println("Relational operator: ", one, two, three)

	// Logical
	wrong := false
	wrongReverse := !wrong
	fmt.Printf("Logical operator: !wrong \t (%t)\n", wrongReverse)

	// Looping
	// Looping 1
	for i:= 0; i < 3; i++ {
		fmt.Println("Angka ", i)
	}

	// Looping 2
	i := 0
	for i < 3 {
		fmt.Println("Angka ", i)
		i++
	}

	// Do while
	j:= 0
	for {
		fmt.Println("Angka ", j)
		j++
		if j == 3 {
			break
		}
	}

	// For range loop
	var dataa []string
	dataa = []string{"aaa", "bbb"}
	for key, data := range dataa {
		fmt.Println("No", key+1, ":", data)
	}
}