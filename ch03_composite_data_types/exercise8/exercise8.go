package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Make changes to csvData.go in order to separate the fields of a record with a character that is given as a command-line argument.

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// CSV file read all at once
	// lines data type is [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func saveCSVFile(filepath string, delimiter rune) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = delimiter
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("csvData input output delimiter!")
		return
	}
	input := os.Args[1]
	output := os.Args[2]
	delimiter := os.Args[3]
	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	// CSV data is read in columns - each line is a slice
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Println(temp)
	}
	err = saveCSVFile(output, []rune(delimiter)[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// string to []rune conversion
	sample := "⌘こ🦡んに🦍ちは"
	var rs []rune = []rune(sample)
	for k := range rs {
		fmt.Printf("%c", rs[k])
	}
	fmt.Println()
}
