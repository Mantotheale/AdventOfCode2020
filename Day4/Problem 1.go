package main

import (
	"fmt"
	"bufio"
	"os"
  "strings"
)

type record struct {
  byr string
  iyr string
  eyr string
  hgt string
  hcl string
  ecl string
  pid string
}

func main() {
  //-------------
  // Slice of records to store the various data in input
  var records []record = ExtractRecords()
  // Var to count the total of valid passports
	var total int
	
  // For each record, if it's valid increase the counter
	for _, r := range records {
		if IsRecordValid(r) {
			total++
		}
	}
	
	fmt.Println(total)
}

// Function to extract the various records from the input file
func ExtractRecords() (records []record) {
  // Open the input file and create a scanner to parse it
  file, _ := os.Open("Input.txt")
	scanner := bufio.NewScanner(file)

  // Var to store the various fields which represents a record
  var fields string
	
	for scanner.Scan() {
    // If a blank row isn't scan, add the row to the record representation
		if scanner.Text() != "" {
			fields += scanner.Text() + " "
		} else {
      // If a blank row is scan, convert the record representation to an actual record and add it to the records slice
			records = append(records, ConvertRecord(fields))

      // Then reset the representation
			fields = ""
		}
	}
  // Add the last record, which is lost due to the end of file
	records = append(records, ConvertRecord(fields))
	file.Close()

  return
}

// Function to convert a string of fields to an actual record
func ConvertRecord(s string) (r record) {
  // Drop the last char pf the string (it's a ' ')
  s = s[:len(s) - 1]

  // Split the string in the various fields
  var fields []string = strings.Split(s, " ")

  // Update the fields of the record based on the various fields
  for _, f := range fields {
    switch f[:3] {
      case "byr": r.byr = f[4:]
      case "iyr": r.iyr = f[4:]
      case "eyr": r.eyr = f[4:]
      case "hgt": r.hgt = f[4:]
      case "hcl": r.hcl = f[4:]
      case "ecl": r.ecl = f[4:]
      case "pid": r.pid = f[4:]
    }
  }

  return
}

// Function to check if a record is valid
func IsRecordValid (r record) bool {
  // If neither of the mandatory fields are missing, the record is valid
  if r.byr != "" && r.iyr != "" && r.eyr != "" && r.hgt != "" && r.hcl != "" && r.ecl != "" && r.pid != "" {
    return true
  } else {
    return false
  }
}
