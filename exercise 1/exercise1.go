package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func shuffle(records [][]string) [][]string {
	rand.Seed(time.Now().UTC().UnixNano())
	//Randomizes the array made from the csv file so that the questions aren't
	//exactly the same every run
	//An inside out method of shuffling a slice/array
	for i := range records {
		j := rand.Intn(i + 1)
		records[i], records[j] = records[j], records[i]
	}
	return records
}

func main() {

	iptr := flag.Int("timer", 5, "Time you get for each question, default is 5 s.")
	bptr := flag.Bool("shuffle", false, "Should the csv file be shuffled?")

	//need to parse the flag variables --spent way too much time trying to figure this out
	flag.Parse()
	//reading the csv file
	//content will be of type []byte
	content, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	//stringifies the content from type []byte
	s := string(content)

	//reads the stringified file
	//records is an array of array of strings [][]string
	r := csv.NewReader(strings.NewReader(s))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *bptr {
		records = shuffle(records)
	}
	fmt.Println("ip: ", *iptr)

	var score int
	var temp string
	fmt.Println("Press the Enter key to start the game")
	fmt.Scanln(&temp)
	for i := range records {
		fmt.Printf("%s= ", records[i][0])
		fmt.Scan(&temp)
		if score == len(records) {
			fmt.Println("You Won the game!")
			break
		}
		if temp == records[i][1] {
			//fmt.Println("Correct!")
			score++
		} else {
			//fmt.Println("Incorrect! Try again later")
			//break
			continue
		}
	}
	fmt.Printf("You got %d question(s) correct out of %d total", score, len(records))
}
