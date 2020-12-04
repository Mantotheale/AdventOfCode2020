package main

import (
	"fmt"
	"bufio"
	"os"
  "strings"
  "strconv"
  "unicode"
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
  // Check all the fields, if all of them are valid, the row is valid
  if IsByrValid(r.byr) && IsIyrValid(r.iyr) && IsEyrValid(r.eyr) && IsHgtValid(r.hgt) && IsHclValid(r.hcl) && IsEclValid(r.ecl) && IsPidValid(r.pid) {
    return true
  } else {
    return false
  }
}

func IsByrValid(byr string) bool {
  if byr == "" {
    return false
  }
  // Convert the field to it's respective number
  value, err := strconv.Atoi(byr)
  
  if err == nil {
    // The byr should be between 1920 and 2002
    if value >= 1920 && value <= 2002 {
      return true
    } else {
      return false
    }
  } else {
    return false
  }
}

func IsIyrValid(iyr string) bool {
  if iyr == "" {
    return false
  }
  // Convert the field to it's respective number
  value, err := strconv.Atoi(iyr)

  if err == nil {
    // The iyr should be between 2010 and 2020
    if value >= 2010 && value <= 2020 {
      return true
    } else {
      return false
    }
  } else {
    return false
  }
}

func IsEyrValid(eyr string) bool {
  if eyr == "" {
    return false
  }
  // Convert the field to it's respective number
  value, err := strconv.Atoi(eyr)

  if err == nil {
    // The eyr should be between 2020 and 2030
    if value >= 2020 && value <= 2030 {
      return true
    } else {
      return false
    }
  } else {
    return false
  }
}

func IsHgtValid(hgt string) bool {
  if hgt == "" {
    return false
  }
  // Split the field in the numerical value and the unit
  var heightString, unit string
  var height int

  for i, c := range hgt {
    if unicode.IsDigit(c) {
      heightString += string(c)
    } else {
      hgt = hgt[i:]
      break
    }
  }
  height, err := strconv.Atoi(heightString)
  unit = hgt

  if err == nil {
    if unit == "cm" {
      // The height in cm should be between 150 and 193
      if height >= 150 && height <= 193 {
        return true
      } else {
        return false
      }
    } else if unit == "in" {
      // The height in inches should be between 59 and 76
      if height >= 59 && height <= 76 {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  } else {
    return false
  }
}

func IsHclValid(hcl string) bool {
  // hcl should be #xxxxxx, so the length should be 7
  if len(hcl) != 7 {
    return false
  }
  // Remove the # from the string
  hcl = hcl[1:]

  // Check if every element of hcl is a digit or a char from 'a' to 'f'
  for _, c := range hcl {
    if !unicode.IsDigit(c) && !(c >= 'a' && c <= 'f') {
      return false
    }
  }

  return true
}

func IsEclValid(ecl string) bool {
  // ecl should be one of those colors
  var validColors [7]string = [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

  for _, c := range validColors {
    if ecl == c {
      return true
    }
  }

  return false
}

func IsPidValid(pid string) bool {
  // pid should be a 9 digits number
  if len(pid) == 9 {
    for _, c := range pid {
      if !unicode.IsDigit(c) {
        return false
      }
    }

    return true
  } else {
    return false
  }
}
