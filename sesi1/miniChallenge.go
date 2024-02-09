package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func miniTest() {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number: ")
	stringNum, _ := input.ReadString('\n')
	stringNum = strings.Replace(stringNum, "\n", "", -1)

	n, err := strconv.Atoi(stringNum)

	if err != nil || n < 1 {
		fmt.Println("Please enter a valid number bigger than 1")
	}

	var i float64 = 1
	for i <= float64(n) {
		cubeRoot := math.Cbrt(i)
		isCube := cubeRoot == math.Trunc(cubeRoot) 
		squareRoot := math.Sqrt(i)
		isSquare := squareRoot == math.Trunc(squareRoot)

		if isCube && isSquare  {
			fmt.Println("SquareCube")
		} else if isCube {
			fmt.Println("Cube")
		} else if isSquare {
			fmt.Println("Square")
		} else {
			fmt.Println(i)
		}

		i++
	}
}