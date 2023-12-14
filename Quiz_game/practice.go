package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	text = formatText(text)
	fmt.Println(text)

	comparison_text := strings.ToLower("hey There")
	fmt.Println(comparison_text)

	if text == comparison_text {
		fmt.Println("this is correct")
	} else {
		fmt.Println("They are not the same")
	} 
}



func formatText(s string) string {
	s = strings.ToLower(s)
    return strings.Join(strings.Fields(s), " ")
}

