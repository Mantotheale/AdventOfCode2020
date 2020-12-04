package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main () {
  // Read all the numbers from the file and save them in a slice
  var number []int = ReadNumbers()

	for i := 0; i < len(numbers) - 1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == 2020 {
				fmt.Println(numbers[i] * numbers[j])
			}
		}
	}
}

func ReadNumbers() (numbers []int) {
  // Create a scanner to parse the file
  scanner := bufio.NewScanner(os.Stdin)
  
  for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())

		if err == nil {
			numbers = append(numbers, n)
		} else {
			fmt.Println("Errore")
		}
	}
}
