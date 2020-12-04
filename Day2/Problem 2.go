package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main() {
	// Counter to the total of valid rows
	var totalValid int

  // Open the input file and create e scanner to parse it
	file, _ := os.Open("Input.txt")
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
    // Vars the two indexes in which check the char, the char specified and the string provided
    var index1, index2 int
    var s string
    var c rune

    // Extract the values from the row of the file
		index1, index2, c, s = ExtractValues(scanner.Text())
		
    // If the row is valid, increment the counter
		if IsRowValid (s, c, index1, index2) {
			totalValid++
		}
	}
	
	fmt.Println(totalValid)
}

// Function to extract the various fields form a row in input
func ExtractValues (row string) (index1 int, index2 int, c rune, s string) {
	// Strings to store the various fields of the string
	var index1String, index2String string
	
  // Extract the first field of the first index
  var i int
	for i = 0; row[i] != '-'; i++ {
		index1String += string(row[i])
	}
	row = row[i + 1:]
	
  // Extract the second field of the second index
	for i = 0; row[i] != ' '; i++ {
		index2String += string(row[i])
	}
	row = row[i + 1:]
	
  // Extract the char to find
	c = rune(row[0])
	row = row[3:]
	
  // Extract the string in which find the char
	s = row
	
  // Convert the fields of the two indexes from strings to int
  index1, _ = strconv.Atoi(index1String)
  index2, _ = strconv.Atoi(index2String)

	return
}

// Function to verify if a row is valid
func IsRowValid (s string, c rune, index1 int, index2 int) bool {
  // Store the two chars to check
  var c1, c2 rune = rune(s[index1 - 1]), rune(s[index2 - 1])

  // If only one of the two chars matches, the row is valid
	if (c1 == c && c2 != c) || (c1 != c && c2 == c) {
		return true
	} else {
		return false
	}
}
