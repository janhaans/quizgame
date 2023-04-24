package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	for _, problem := range problems {
		fmt.Printf("Problem = %s, Solution = %s\n", problem[0], problem[1])
	}

}
