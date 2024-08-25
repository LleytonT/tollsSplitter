package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type TollRecord struct {
	Date string
	Type string
	Details string
	Concession string
	Credit string
	Balance string
}

func main() {
	file, err := os.Open("tolls.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var records []TollRecord
    for _, line := range lines {
        record := TollRecord{
            Date:       line[0],
            Type:       line[1],
            Details:    line[2],
            Concession: line[3],
            Credit:     line[4],
            Balance:    line[5],
        }
        records = append(records, record)
    }

		totalCredit := 0.0
    for _, record := range records {
				credit := record.Credit
				if credit != "Debit/Credit amount" {
					creditValue, err := strconv.ParseFloat(strings.ReplaceAll(credit, "-", ""), 3)
					if err != nil {
						log.Fatal(err)
					}
					totalCredit += creditValue
				}
        fmt.Printf("%+v\n", record)
    }

		fmt.Println("Total Tolls: ", totalCredit)
}