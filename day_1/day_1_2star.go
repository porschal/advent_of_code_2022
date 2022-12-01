package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	sum := 0

	content, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	var calories []int

	for scanner.Scan() {

		if scanner.Text() != "" {
			i, _ := strconv.Atoi(scanner.Text())
			sum += i
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}
	// fmt.Print(max)
	sort.Ints(calories)

	calories = calories[len(calories)-3:]
	sum = 0
	for _, v := range calories {
		sum += v
	}

	fmt.Println(sum)

}
