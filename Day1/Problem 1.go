package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main () {
  // Read all the numbers from the file and save them in a slice
  var numbers []int = ReadNumbers()

  // Scan every combination of numbers in the slice
	for i := 0; i < len(numbers) - 1; i++ {
		for j := i + 1; j < len(numbers); j++ {
      // Check if the sum is 2020
			if numbers[i] + numbers[j] == 2020 {
        // Print the product of the numbers if found
				fmt.Println(numbers[i] * numbers[j])
			}
		}
	}
}

func ReadNumbers() (numbers []int) {
  // Open the input file and create a scanner to parse it
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)
  
  for scanner.Scan() {
    // Try to convert the row to the integer in it
    n, err := strconv.Atoi(scanner.Text())

    // If the conversion is successful, add the number to the slice
    if err == nil {
      numbers = append(numbers, n)
    } else {
      fmt.Println("Errore")
    }
  }

  return
}
