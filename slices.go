package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	names := []string{"Thomas", "Agumon", "Lala", "Bare"}
	values := []int{11, 12, 13, 14}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Agumon"))
	fmt.Println(slices.Index(names, "Lala"))
	fmt.Println(slices.Index(names, "Bare"))

}
