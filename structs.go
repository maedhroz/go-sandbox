package main

import "fmt"

type SuperGuy struct {
	superValue string
}

type SubGuy struct {
	SuperGuy
	value string
}



func main() {
	i := 4000000000
	f := float64(i)
	u := uint64(i)

	fmt.Println(i, f, u)

	var thingy = &SubGuy{}
	thingy.superValue = "asfdfds"

	fmt.Println(thingy.superValue)

	a := []string {"Caleb", "Rackliffe"}

	fmt.Println(len(a))

	if nil {

	}
}
