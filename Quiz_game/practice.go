package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	val := [][]int{{1,2},{3,4},{5,6}}

// 	fmt.Println(Shuffle(val))
//   }

// func Shuffle(vals [][]int) [][]int {
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	for n := len(vals); n > 0; n-- {
// 		randIndex := r.Intn(n)
// 		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
// 	}
// 	return vals
// }