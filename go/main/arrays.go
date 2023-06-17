package main

import "fmt"

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	// friends[1] will change the value of friends[1] in the array
	// because of the &friends in the range loop
	for i, v := range &friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("Aft[%s]", v)
		}
	}

}
