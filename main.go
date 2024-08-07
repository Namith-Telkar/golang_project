package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
	fmt.Println("Bye World!")
	fmt.Println(":D")

	name := "Namith"
	age := 23

	fmt.Println("My name is", name, "and I am", age, "years old.")
	
	fmt.Printf("My name is %v and I am %v years old.\n", name, age)
	fmt.Printf("My name is %q and I am %vqyears old.\n", name, age)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("You scored %0.1f points!\n", 225.55)

	var str string = fmt.Sprintf("My name is %v and I am %v years old.\n", name, age)
	fmt.Printf("The saved string is: %v", str)


}