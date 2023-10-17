package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//Presents quiz problem, accept user input, check for correctness

func main(){
	csvFilename := flag.String("csv", "problems.csv","a csv file in the format of 'question, answer' ")

	timeLimit := flag.Int("limit", 30, "this is the time limit fo rthe quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)


	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	// <-timer.C //here we are waiting for a message from a channel so our code will block until we get that message 

	//print out the problems, get a response from the user then check if the user is correct 
	// we add a select statement to the for loop that tells us if we get a message from the timer channel then we know that we need to stop the for loop and stop presenting problems 
	correct := 0
	
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1,p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer) 
			answerCh <- answer
		}()

		//so what this select statement means is that, if we get a message back from the answer channel first, then the timer hasnt run out but if we get from the timer channel first then the answer isnt valid and we return to end the select statement 
		select { //when we get a message from the timer channel, the for loop will stop, we dont use a break in this situation because it only breaks out of the select and not the for loop
			// the return statement breaks out of both
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh: //if we get an answer from the answer channel , if check to see if it is correct 
			if answer == p.a {
				correct++
			}
		}
		
	}
	
}


// this takes in a 2D string slice and returns a slice of problem struct
func parseLines(lines [][]string) []problem{
	//first we declare the variable we are going to return
	// the size is len(lines) because that gives us the total number of rows in our csv file 
	// ret := make([]problem, len(lines))
	// for i, line := range lines {
	// 	ret[i] = problem{
	// 		q: line[0],
	// 		a: line[1],
	// 	}
	// }

	// return ret 

	//using the append method to always resize the slice
	ret := make([]problem, 0)//initialise empty slice 
					
	for _, line := range lines{
		ret = append(ret, problem{q: line[0], a: strings.TrimSpace(line[1])})
	}

	return ret

}


type problem struct {
	q string 
	a string
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

