package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	content, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	s := string(content)

	r := csv.NewReader(strings.NewReader(s))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//Randomizes the array made from the csv file so that the questions aren't
	//exactly the same every run
	//An inside out method of shuffling a slice/array
	for i := range records {
		j := rand.Intn(i + 1)
		records[i], records[j] = records[j], records[i]
	}

	var score int
	var temp string
	for i := range records {
		fmt.Printf("%s= ", records[i][0])
		fmt.Scan(&temp)
		if score == len(records) {
			fmt.Println("You Won the game!")
			break
		}
		if temp == records[i][1] {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect! Try again later")
			break
		}
	}
	fmt.Printf("You got %d question(s) correct out of %d total", score, len(records))
}
