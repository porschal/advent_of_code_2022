package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//  A ROCK 1
// B Paper 2
// C scissors 3

// X loose
// Y  draw
// Z   win

// 0 LOSE
// 3 DRAW
//6 WON

func main() {
	// sum := 0

	content, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)
	sum := 0

	for scanner.Scan() {
		str := scanner.Text()
		before, after, _ := strings.Cut(str, " ")

		if before == "A" { // Rock
			if after == "X" { // wIN
				sum += 3 //draw + rock
			} else if after == "Y" { // draw
				sum += 4 // win + pappers
			} else { //lose
				sum += 8 // lose + scissors
			}
			continue
		}

		if before == "B" { // paper
			if after == "X" { //rock
				sum += 1
			} else if after == "Y" { // paper
				sum += 5
			} else { //scissors
				sum += 9
			}
			continue
		}

		if before == "C" { //scissors
			if after == "X" { //rock
				sum += 2
			} else if after == "Y" {
				sum += 6
			} else {
				sum += 7
			}
			continue
		}
	}

	fmt.Print(sum)
}
