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
    // Vars to save max occurencies, min occurencies, the actual occurencies, the char specified and the string provided
    var minOcc, maxOcc, occ int
    var s string
    var c rune

    // Extract the values from the row of the file
		minOcc, maxOcc, c, s = ExtractValues(scanner.Text())

    // Count the occurencies of the character c in the string s
		occ = CountOccurencies(s, c)
		
    // If the occurencies are between the bounds, increase the counter
		if occ >= minOcc && occ <= maxOcc {
			totalValid++
		}
	}
	
	fmt.Println(totalValid)
}

// Function to extract the various fields form a row in input
func ExtractValues (row string) (minOcc int, maxOcc int, c rune, s string) {
  // Strings to store the various fields of the string
	var minOccString, maxOccString string
	
  // Extract the first field of the minimum occurencies
  var i int
	for i = 0; row[i] != '-'; i++ {
		minOccString += string(row[i])
	}
	row = row[i + 1:]
	
  // Extract the second field of the maximum occurencies
	for i = 0; row[i] != ' '; i++ {
		maxOccString += string(row[i])
	}
	row = row[i + 1:]
	
  // Extract the char to find
	c = rune(row[0])
	row = row[3:]
	
  // Extract the string in which find the char
	s = row
	
  // Convert the fields of max e min occurencies from strings to int
  minOcc, _ = strconv.Atoi(minOccString)
  maxOcc, _ = strconv.Atoi(maxOccString)

	return
}

// Function to count the occurencies of a given char in a string
func CountOccurencies (s string, c rune) (occ int) {
  // Scan every element of the string
	for _, v := range s {
    // If found, increment the counter
		if v == c {
			occ++
		}
	}
	
	return
}
