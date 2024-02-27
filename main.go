//This is the example pointer value

package main

import (
	"fmt"
)

func update(a *int) {
	fmt.Println(*a) //dereferencing and print
	*a = *a + 10    //Assign pointer value and deferencing a variable  + new value
}

func main() {

	x := 10 //assignment

	y := &x //referencing

	fmt.Println(y) // Accessing pointer variable value

	update(&x) //Accessing Memory Address

	println(x) // Accessing

}
