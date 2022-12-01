package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum := 0
	max := 0

	content, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		if scanner.Text() != "" {
			i, _ := strconv.Atoi(scanner.Text())
			sum += i
		} else {
			if sum >= max {
				max = sum
			}
			sum = 0
		}
	}

	fmt.Print(max)

}
