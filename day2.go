package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
	"math"
	"slices"
)

func main() {
	// Open the file
	file, err := os.Open("inputd2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

  safe := 0 //number of safe reports

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a list to store the lines
	var lines []string

	// Read each line and append it to the list
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Iterate over the list
  for index, _ := range lines {
      if isSafe(lines[index]) {
        safe = safe + 1
      } else { //problem dampener
        numbers := strings.Split(lines[index], " ")
          for i:=0; i < len(numbers); i++ {
            if isSafe(strings.Join(RemoveIndex(numbers, i), " ")) {
              safe = safe + 1
              break
            }
          }
      }
  }

  fmt.Println("The number of safe reports is: ", safe)
}

func RemoveIndex(s []string, index int) []string {
    ret := make([]string, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func isSafe(line string) bool {
  retVal := true

  // Split the string into an array of strings
	numbers := strings.Split(line, " ")
  intSlice := make([]int, 0)

	//Iterate over the array and convert each string to an integer
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue // Skip to the next iteration
		}

		// Process the number
		// fmt.Println(intSlice)

		intSlice = append(intSlice, num)
	}

  // 1. sort the slice ascending and descending and compare
  // to orginal slice. If neither are equal then fail.
  // 2. iterate over the original slice and compare current
  // versus previous values. If difference is not 1,2, or 3 then fail.
  // 3. if 1 & 2 pass then pass/true.

  ascending := make([]int, len(intSlice))
  copy(ascending, intSlice)
  descending := make([]int, len(intSlice))
  copy(descending, intSlice)


  sort.Ints(ascending)
  sort.Slice(descending, func(i, j int) bool {
      return descending[j] < descending[i]
   })



   if slices.Equal(intSlice, ascending) ||  slices.Equal(intSlice, descending) {
     for index:=0; index < len(ascending); index++ {
        if index < (len(ascending)-1) {
          diff := math.Abs(float64(ascending[index] - ascending[index+1]))
          if !(diff == 1 || diff == 2 || diff == 3) {
            retVal = false
          }
        }
     }
   } else {
     retVal = false
   }

	return retVal
}






