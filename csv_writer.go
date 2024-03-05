package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Thomas", "Tommy", "Tomtom"})
	_ = writer.Write([]string{"Rendra", "Bana", "Bare"})
	_ = writer.Write([]string{"Patamon", "Agumon", "Balmon"})

	writer.Flush()

}
