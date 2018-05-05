package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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

	fileptr := flag.String("csv", "problems.csv", "filename of the csvfile(default is problems.csv)")
	timeptr := flag.Int("timer", 5, "Time you get for each question, default is 5 s.")
	shuffleptr := flag.Bool("shuffle", false, "Should the csv file be shuffled?")

	//need to parse the flag variables --spent way too much time trying to figure this out
	flag.Parse()
	//reading the csv file
	//content will be of type []byte
	content, err := ioutil.ReadFile(*fileptr)
	if err != nil {
		fmt.Printf("Unable to open file: %s\n", *fileptr)
		os.Exit(1)
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

	if *shuffleptr {
		records = shuffle(records)
	}
	fmt.Println("ip: ", *timeptr)

	var score int
	var temp string
	fmt.Println("Press the Enter key to start the game")
	fmt.Scanln(&temp)
	for i := range records {
		timer := time.NewTimer(time.Second * time.Duration(*timeptr))

		fmt.Printf("%s= ", records[i][0])
		answerCh := make(chan string)
		// if *scoreptr == len(records) {
		// 	fmt.Println("You Won the game!")
		// }
		go func() {
			fmt.Scan(&temp)
			answerCh <- temp
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime ranout")
			fmt.Printf("You got %d question(s) correct out of %d total", score, len(records))
			return
		case answer := <-answerCh:
			if answer == records[i][1] {
				score++
			}
		}

	}
	fmt.Printf("You got %d question(s) correct out of %d total", score, len(records))

}
