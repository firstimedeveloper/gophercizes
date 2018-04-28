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
	//numOfQ := len(records)
	var score int
	var temp string
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		randNum := seed.Intn(len(records))

		fmt.Printf("%s= ", records[randNum][0])
		fmt.Scan(&temp)
		// if score == len(records) {
		// 	fmt.Println("You Won the game!")
		// 	break
		//}
		if temp == records[randNum][1] {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect! Try again later")
			break
		}
	}
	fmt.Printf("You got %d questions correct!", score)
}
