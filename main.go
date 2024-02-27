//Now we are learning array Slice

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// We are learning array Slice

	// An Array is a composite data type or data structure

	var students [3]string

	students[0] = "Tahmid"
	students[1] = "Aksa"
	students[2] = "Abud Ubaida"

	sl := students[0:2]

	fmt.Println(sl)

	x := make([]float32, 5)
	x[0] = 1
	x[1] = 2

	fmt.Println(x)

	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", students)

	var fruits []string

	fruits = append(fruits, "Apple", "Banana", "Mango")

	test := reflect.TypeOf(fruits).Kind().String()

	fmt.Println(test)

	fmt.Println(fruits)

}
