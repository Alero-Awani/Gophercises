package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// func main(){
// 	f, err := os.Open("problems.csv")
// 	if err != nil{
// 		exit("Failed to open the csv file")
// 	}
// 	reader := csv.NewReader(f)

// 	records, err := reader.ReadAll()

// 	if err != nil {
// 		exit("There was an error when reading the records")
// 	}

// 	fmt.Println(Shuffle(records))
// }

// func exit(msg string) {
// 		fmt.Println(msg)
// 		os.Exit(1)
// 	}

// func Shuffle(vals [][]string) [][]string {
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	for n := len(vals); n > 0; n-- {
// 		randIndex := r.Intn(n)
// 		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
// 	}
// 	return vals
// }

//readall returns a slice of slice, so we want to shuffle the slices inside the slice before printing it out



