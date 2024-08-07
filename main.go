package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"A", "B", "C", "D"}

	names[1] = "E"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores = []float64{10.2, 20.5, 30.3}
	scores[1] = 25.5
	scores = append(scores, 85.5)
	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]	
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}