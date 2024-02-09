package main

import "fmt"

func main() {

	//Slice
	animals1 := []string{"cat", "dog"}
	animals2 := []string{"horse", "duck"}

	// Append function
	animals := append(animals1, animals2...)
	fmt.Println("Append function: ", animals)

	// Copy function
	copyFunc := copy(animals1, animals2)
	fmt.Println("Animals1: ", animals1)
	fmt.Println("Animals2: ", animals2)
	fmt.Println("Copied elements: ", copyFunc)

	// Slicing function
	fruits := []string{"apple", "mango", "orange", "pear", "banana"}

	var fruits1 = fruits[2:4]
	var fruits2 = fruits[:4]
	var fruits3 = fruits[1:]
	var fruits4 = fruits[:] // equal to fruits[:len(fruits)]

	fmt.Println("Fruits1: ", fruits1)
	fmt.Println("Fruits2: ", fruits2)
	fmt.Println("Fruits3: ", fruits3)
	fmt.Println("Fruits4: ", fruits4)

	// Slice (backing array)
	lang1 := []string{"python", "ruby", "java", "go", "js"}
	lang2 := lang1[2:4]
	lang2[0] = "rust"
	fmt.Println("Lang1: ", lang1)
	fmt.Println("Lang2: ", lang2) // lang1 value will change too

	// Slice cap function
	fmt.Println("--- Slice cap function ---")
	newFruits := []string{"apple", "mango", "durian", "banana"}
	fmt.Println("Fruits cap:", cap(newFruits)) // 4
	fmt.Println("Fruits len:", len(newFruits)) // 4

	newFruits2 := newFruits[0:3]
	fmt.Println("Fruits2 cap:", cap(newFruits2)) // 4
	fmt.Println("Fruits2 len:", len(newFruits2)) // 3

	newFruits3 := newFruits[1:]
	fmt.Println("Fruits3 cap:", cap(newFruits3)) // 3 because starting from index 1
	fmt.Println("Fruits3 len:", len(newFruits3)) // 3

	// Map
	fmt.Println("--- Map ---")
	// -- Initialize 1
	var person1 map[string]string // declaration
	person1 = map[string]string{} // initialization

	person1["name"] = "Jennie"
	person1["age"] = "27"

	fmt.Println("Person1: ", person1)

	// -- Initialize 2
	var person2 = map[string]string{
		"name": "Lisa",
		"age": "26",
	}

	fmt.Println("Person2: ", person2)

	delete(person2, "age")
	fmt.Println("Person2 after deletion: ", person2)

	value, isExisting := person2["name"]
	fmt.Println("value, isExisting: ", value, isExisting)
	
  value, isExisting = person2["test"]
  fmt.Println("value, isExisting: ", value, isExisting) // isExisting will be false

	// Combine slice with map
	people := []map[string]string{
		{"name": "Jennie", "age": "26"},
		{"name": "Rose", "age": "27"},
		{"name": "Jisoo", "age": "28"},
		{"name": "Lisa", "age": "29"},
	}

	for key, person := range people {
		fmt.Println(key, person["name"])
	}

	// Pointer
	fmt.Println("--- Pointer ---")
	// Kapan kita pakai pointer?
	// - Menghindari penggandaan data
	// - Nilai kosong yg bermakna
	// - Mengubah nilai variable di luar fungsi

	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value): ", firstNumber)
	fmt.Println("firstNumber (memory address): ", &firstNumber)
	fmt.Println("secondNumber (value): ", *secondNumber)
	fmt.Println("secondNumber (memory address): ", secondNumber)

	*secondNumber = 8
	fmt.Println("firstNumber (value): ", firstNumber)
	fmt.Println("secondNumber (value): ", *secondNumber)

	// Pointer as parameter
	fmt.Println("--- Pointer as parameter ---")
	newNumber := 10
	fmt.Println("Old no: ", newNumber)
	changeValue(&newNumber)
	fmt.Println("New no: ", newNumber)
}

func changeValue (number *int) {
	*number = 20
}