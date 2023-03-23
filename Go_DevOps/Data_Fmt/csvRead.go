package main

import (
	"fmt"
	"strings"
)

// fakeContent represents a file's content.
const fakeContent = `
John,Doak
Sarah, Murphy
David, Justice
`

// record represents a record containing first and last names that are
// stored in a csv.
type record []string

// validate validates if the csv line had the correct number of entries.
func (r record) validate() error {
	if len(r) != 2 {
		return fmt.Errorf("data format is incorrect")
	}
	return nil
}

// first  and last returns the record's first aand last name.
func (r record) first() string {
	return r[0]
}
func (r record) last() string {
	return r[1]
}

// readRecs reads a file in csv format representing records. It returns the records.
// This will skip any lines that are blank and stops on the first error encountered.
func readRecs() ([]record, error) {
	// In this example we just take from fakeContent instead of
	// opening a file.

	lines := strings.Split(fakeContent, "\n") // Split by line
	var records []record
	for i, line := range lines {
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		var rec record = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", i, err)
		}
		records = append(records, rec)
	}
	return records, nil
}

func main() {
	recs, err := readRecs()
	if err != nil {
		panic(err)
	}
	for _, rec := range recs {
		fmt.Println(rec.first())
	}
}


// If the file is large and we want to be efficient, we can use the bufio and bytes packages:

// func readRecs() ([]record, error) { 
//     file, err := os.Open("data.csv") 
//     if err != nil { 
//         return nil, err 
//     } 
//     defer file.Close() 
//     scanner := bufio.NewScanner(fakeFile) 
//     var records []record 
//     lineNum := 0 
//     for scanner.Scan() { 
//         line := scanner.Text() 
//         if strings.TrimSpace(line) == "" { 
//             continue 
//         } 
//         var rec record = strings.Split(line, ",") 
//         if err := rec.validate(); err != nil { 
//             return nil, fmt.Errorf("entry at line %d was invalid: %w", lineNum, err) 
//         } 
//         records = append(records, rec) 
//         lineNum++ 
//     } 
// return records, scanner.Err()
// }

// func main() {
// 	// Create our fake file
// 	fakeFile.WriteString(fakeContent)
	

// 	recs, err := readRecs()
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, rec := range recs {
// 		fmt.Println(rec.first())
// 	}
// }