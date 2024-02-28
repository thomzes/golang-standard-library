package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Thomas Ardiansah", "Thom"))
	fmt.Println(strings.Split("Thomas Ardiansah", " "))
	fmt.Println(strings.ToLower("Thomas Ardiansah"))
	fmt.Println(strings.ToUpper("thomas ardiansah"))
	fmt.Println(strings.Trim("      Thomas Ardiansah        ", " "))
	fmt.Println(strings.ReplaceAll("Thomas Ardiansah Pambudi", "Thomas", "Ardiansah"))
}