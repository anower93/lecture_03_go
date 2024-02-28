//Now we are learning Struct /Composite Data type

package main

import (
	"fmt"
	
)

func main() {


type Book struct {
	name string
	price float32
	pages int
}


var b1 Book 

b1.name = "Go in Bangla"
b1.price = 13.13
b1.pages = 100

fmt.Println(b1)
fmt.Println(b1.name)
fmt.Println(b1.pages)




}
