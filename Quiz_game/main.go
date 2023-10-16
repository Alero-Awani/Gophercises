package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

//Presents quiz problem, accept user input, check for correctness

func main(){
	csvFilename := flag.String("csv", "problems.csv","a csv file in the format of 'question, answer' ")
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

	//print out the problems, get a response from the user then check if the user is correct 
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1,p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) //using the reference to answer here to make sure that it has a pointer value to work with so whenever it sets the value we can access it with our variable 
		
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	
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

