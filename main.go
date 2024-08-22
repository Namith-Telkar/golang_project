package main

import (
	"fmt"
)

func main() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }
	
	// for i:=0; i<5; i++ {
	// 	fmt.Println("The Value of i is:", i)
	// }

	names := []string{"John", "Robb", "Sansa", "Arya", "Bran"}

	// for i:=0; i<len(names); i++{
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names{
	// 	fmt.Printf("The value at index %v is %v\n", index, value)
	// }

	for _, value := range names{
		fmt.Printf("The value is %v \n", value)
	}
}