package tsvreader

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadRandomLines(filename string, n int) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	allRecords, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Extract values from the first and third columns
	extractedRecords := make([][]string, len(allRecords))
	for i, record := range allRecords {
		if len(record) >= 4 { // Check to ensure the record has at least 4 columns
			extractedRecords[i] = []string{record[1], record[3]}
		} else {
			// Handle the case where there are less than 3 columns (based on your requirements)
			return nil, fmt.Errorf("record on line %d has less than 3 columns", i+1)
		}
	}

	if n > len(extractedRecords) {
		return nil, fmt.Errorf("requested more lines than available")
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(extractedRecords), func(i, j int) { extractedRecords[i], extractedRecords[j] = extractedRecords[j], extractedRecords[i] })

	return extractedRecords[:n], nil
}
