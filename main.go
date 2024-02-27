//Now we are learning array

package main

import (
	"fmt"
)

func main() {

	// We are learning array

	var students [3]string

	students[0] = "Tahmid"
	students[1] = "Aksa"
	students[2] = "Abu Ubaida"

	fmt.Println(students) // In this way we can declare array and pull data, retrieve data or get data

	// Or we can write in short hand

	teachers := [3]string{"Tahmid", "Abu Ubaida", "Aksa"}

	fmt.Println(teachers)

	//We can import unlimited data implicitly

	ssc := [...]string{"Aman", "Zaman", "Rahman", "Lalcan"}

	fmt.Println(ssc)

}
