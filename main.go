package main

import (
	"fmt"
	"strconv"
)

var (
	actorName = "James Bond"
	companion = "Sarah Smith"
	age       = 17
	i         = 5
)

func main() {
	i = 30
	j := 42.5
	fmt.Println(i + age)
	fmt.Println(actorName + " " + companion)
	fmt.Printf("%v, %T\n", j, j)

	var k float32
	k = float32(i)
	fmt.Printf("%v, %T\n", k, k)

	var textAge string
	textAge = strconv.Itoa(age)
	fmt.Printf("%v, %T\n", textAge, textAge)
}
