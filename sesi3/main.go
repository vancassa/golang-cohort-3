package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Employee struct {
	name 			string
	age				int
	position	string
}

type Member struct {
	employee 	Employee
	hasAccess bool
}

func (e Employee) introduce(message string) string {
	return fmt.Sprintf("%s, my name is %s and I'm %d years old", message, e.name, e.age)
}

// when you want to change the value of a property, you have to use pointer
func (e *Employee) changeName(newName string) {
	e.name = newName 
}

func main() {
	/////////
	fmt.Println("--- Struct ---")

	var employee Employee
	employee.name = "Fitri"
	employee.age = 20
	employee.position = "Backend"
	fmt.Println(employee)

	var employee2 = Employee{name: "Necha", age: 17, position: "Student"}
	employee2.age = 22
	var employee3 *Employee = &employee2
	fmt.Println("employee2: ", employee2)
	fmt.Println("employee3: ", employee3)

	employee3.name = "Not Necha"
	fmt.Println("employee2: ", employee2)
	fmt.Println("employee3: ", employee3)

	/////////
	fmt.Println("--- Embed struct ---")
	member1 := Member{ employee: Employee{ name: "Member", age: 21, position: "Manager" }, hasAccess: true }
	fmt.Println("Member1: ", member1)

	/////////
	fmt.Println("--- Anonymous struct ---")
	anonStruct := struct{
		name	string
		hobby	string
		rank	int
	}{ 
		name: "Anon", hobby: "Reading manga", rank: 10,
	}
	fmt.Println("Anon: ", anonStruct)


	/////////
	fmt.Println("--- Slice of anonymous struct ---")
	sliceAnon := []struct{
		name 	string
		age 	int
	}{
		{ name: "A", age: 1},
		{ name: "B", age: 2},
		{ name: "C", age: 3},
	}
	fmt.Println("sliceAnon: ", sliceAnon)


	/////////
	fmt.Println("--- Function ---")
	greet("Necha", "ninja", true)
	message := greetReturn("Necha", "ninja", true)
	fmt.Println("Function that returns: ", message)

	multiplication, division := calculate (10, 5)
	fmt.Println("Function returns multiple:", multiplication, division)

	multiple, divide := calculateWithPredefined(10, 5)
	fmt.Println("Pre-defined return value:", multiple, divide)

	fmt.Println("Variadic function", variadicFunction2(1,2,3))
	variadicFunction("Necha", "Twice", "Momusu", "SNSD")

	fmt.Println("--- Method ---")
	fmt.Println(employee.introduce("Yo what's up"))

	employee.changeName("Fitri2")
	fmt.Println(employee.introduce("Yo what's up"))

	/////////
	fmt.Println("--- Reflect ---")
	val := reflect.ValueOf(employee)
	typ := reflect.TypeOf(employee)
	fmt.Println("val", val, "type", typ)

	
}



func greet (name, position string, cond bool) {
	fmt.Println("Hello, my name is", name, "and I'm a", position, ".")
}

func greetReturn (name, position string, cond bool) string {
	return "Hello, my name is "+ name+ " and I'm a"+ position+ "."
}

func calculate (num1, num2 float64) (float64, float64) {
	multiplied := num1 * num2
	divided := num1 / num2
	return multiplied, divided
}

func calculateWithPredefined(num1, num2 float64) (multiplied float64, divided float64) {
	multiplied = num1 * num2
	divided = num1 / num2
	return //still must return, it will return 'multiplied' and 'divided'
}

func variadicFunction (name string, favGroups ...string)  {
	mergeFavGroups := strings.Join(favGroups, ", ")

	fmt.Println("Hello, I'm", name)
	fmt.Println("My favourite groups are", mergeFavGroups)
}

func variadicFunction2 (num ...int) int {
	total := 0

	for _, value := range num {
		total += value
	}

	return total
}