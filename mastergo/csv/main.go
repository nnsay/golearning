package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Record struct {
	Id      int
	Name    string
	Address string
}

var myData = []Record{}

func readCSV(path string) [][]string {
	_, err := os.Stat(path)
	if err != nil {
		log.Panic(err)
	}
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Panic(err)
	}
	return records
}

func saveCSV(path string) {
	csvFile, err := os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	defer csvFile.Close()
	csvWriteer := csv.NewWriter(csvFile)
	csvWriteer.Comma = '\t'
	for _, r := range myData {
		csvWriteer.Write([]string{strconv.Itoa(r.Id), r.Name, r.Address})
	}
	csvWriteer.Flush()
}

// go run mastergo/csv/main.go ./debugfile/source.csv ./debugfile/out2.csv
func main() {
	if len(os.Args) != 3 {
		log.Panic("go run main.go SourceFilePath TargetSoucePath")
	}
	sourceFile := os.Args[1]
	records := readCSV(sourceFile)
	for _, r := range records {
		id, _ := strconv.Atoi(r[0])
		myData = append(myData, Record{id, r[1], r[2]})
	}
	targetFile := os.Args[2]
	saveCSV(targetFile)
}
