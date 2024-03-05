package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`t([a-z])s`)

	fmt.Println(regex.MatchString("tos"))
	fmt.Println(regex.MatchString("Thomass"))
	fmt.Println(regex.MatchString("ThOmas"))

	fmt.Println(regex.FindAllString("tos tis tus tes not thom lala pfft prr", 10))

}
