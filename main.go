//Now we are learning array Map

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// We are learning array Map, It is called Associative array in PHP

	// An Array is a composite data type or data structure

	x := make(map[string]string)

	x["name"] = "Anower"
	x["age"] = "30"
	x["height"] = "5.5"

	//We can delete any item from map

	delete(x, "age")

	fmt.Println(x)

	data := reflect.TypeOf(x).Kind().String()

	fmt.Println(data)

}
