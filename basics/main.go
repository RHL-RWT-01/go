package main

import (
	"fmt"
	"reflect"
)

// greet is a simple function that returns a greeting message
func greet(name string) string {
	return "Hello, " + name
}

// add demonstrates returning multiple values
func add(a, b int) (int, int) {
	return a + b, a * b
}

type Address struct {
	city, state string
}

// Person is a struct (custom data type)
type Person struct {
	name    string
	age     int
	address Address
}

func (p *Person) updateAge() {
	p.age += 1
	fmt.Println("Updated age inside method:", p.age)
}

func main() {
	// 1. Variables and Constants
	var a int = 10 // explicit type
	b := 20        // type inferred
	const pi float64 = 3.1415
	var name string = "Rahul"
	var isGoFun bool = true
	var x, y = 1, 2
	fmt.Println("a:", a, "b:", b, "pi:", pi, "name:", name, "isGoFun:", isGoFun, "x:", x, "y:", y)

	// 2. Built-in Data Types
	var i int = 42
	var f float64 = 3.14
	var s string = "GoLang"
	var bl bool = false
	fmt.Println("Types:", reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(s), reflect.TypeOf(bl))

	// 3. Arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// 4. Slices (dynamic arrays)
	sl := []string{"go", "is", "fun"}
	sl = append(sl, "!")
	fmt.Println("Slice:", sl)

	// 5. Maps (key-value pairs)
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("Map:", m)

	// 6. Structs
	p := Person{name: "Rahul", age: 30, address: Address{city: "Indore", state: "Madhya Pradesh"}}
	fmt.Println("Struct:", p)
	p.updateAge()
	fmt.Println("Age after method call:", p.age)
	fmt.Println("Address after method call:", p.address)

	// 7. Pointers
	var ptr *int = &a
	fmt.Println("Pointer value:", *ptr)

	// 8. Conditionals
	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a equals b")
	} else {
		fmt.Println("a is less than b")
	}

	// 9. Switch statement
	switch name {
	case "Rahul":
		fmt.Println("Hello Rahul!")
	case "Alice":
		fmt.Println("Hello Alice!")
	default:
		fmt.Println("Unknown name")
	}

	// 10. Loops
	// for loop (like while)
	cnt := 1
	for cnt < 5 {
		fmt.Println("cnt:", cnt)
		cnt++
	}

	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	// range loop (over slice)
	for idx, val := range sl {
		fmt.Printf("sl[%d]=%s\n", idx, val)
	}

	// range loop (over map)
	for k, v := range m {
		fmt.Printf("m[%s]=%d\n", k, v)
	}

	// 11. Functions
	fmt.Println(greet("Rahul"))
	sum, prod := add(3, 4)
	fmt.Println("Sum:", sum, "Product:", prod)

	// 12. Zero values
	var zeroInt int
	var zeroStr string
	fmt.Println("Zero values:", zeroInt, zeroStr)

	// 13. Type conversion
	var num int = 100
	var numF float64 = float64(num)
	fmt.Println("Type conversion:", numF)

	// 14. Defer
	defer fmt.Println("This is deferred and prints last!")
	fmt.Println("End of main function.")

	// printMaths()
}
