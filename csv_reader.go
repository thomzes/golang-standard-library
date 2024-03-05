package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "thomas, ardiansah, kurniawan\n" +
		"bana, siti, lia\n" +
		"jeni, lala, popo"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)

	}

}
