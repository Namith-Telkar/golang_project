package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello, world!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "lo"))
	fmt.Println(strings.Split(greeting, ","))

	ages := []int{20, 25, 30, 35, 40, 24, 28, 32, 36, 42}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 80)
	fmt.Println(index)

	names := []string{"John", "Paul", "George", "Ringo"}
	sort.Strings(names)
	fmt.Println(names)

}