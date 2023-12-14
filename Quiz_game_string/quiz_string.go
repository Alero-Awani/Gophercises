package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the format of 'question, answer' ")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to Parse the provided csv file")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)




	// fmt.Println(problems)
	correct := 0 
	for i, p := range problems {
		fmt.Printf("Question %d: %s \n", i+1, p.q)//print out the problem

		answerCh := make(chan string)// create channel
		go func() { // anonymous go routine that takes in answer and sends it to the channel, but why is it a goroutine
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answer := scanner.Text()
			answer = formatText(answer)
			answerCh <- answer
		}()

		
		select {
		case <- timer.C: 
			fmt.Printf("Question %d: %s \n", i+1, p.q)
			return
		case answer := <- answerCh:
			if answer == strings.ToLower(p.a) {
				correct++
			}
			
		}
	}
	fmt.Printf("You scored %d out of %d.\n",correct, len(problems))
}

//changes code to lowercase and removes trailing spaces
func formatText(s string) string {
	s = strings.ToLower(s)
    return strings.Join(strings.Fields(s), " ")
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem , len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}