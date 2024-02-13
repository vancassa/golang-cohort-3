package helpers

import "fmt"

func Greet() {
	fmt.Println("Hello from public function")
}

func greet() {
	fmt.Println("Hello from private function")
}