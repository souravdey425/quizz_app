// package main

// import (
// 	"encoding/csv"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"time"
// )

// type problem struct {
// 	q string
// 	a string
// }

// func problemPuller(fileName string) ([]problem, error) {
// 	fobj, err := os.Open(fileName)
// 	if err != nil {
// 		handleError(err)
// 		return nil, err
// 	}
// 	csvR := csv.NewReader(fobj)
// 	cLines, err := csvR.ReadAll()
// 	if err != nil {
// 		handleError(err)
// 		return nil, err
// 	}
// 	return parseProblem(cLines), nil
// }

// func parseProblem(lines [][]string) []problem {
// 	r := make([]problem, len(lines))
// 	for i := 0; i < len(lines); i++ {
// 		r[i] = problem{q: lines[i][0], a: lines[i][1]}
// 	}
// 	return r
// }

// func handleError(err error) {
// 	fmt.Println(err)

// }
// func exit(msg string) {
// 	fmt.Println(msg)
// 	os.Exit(1)
// }
// func main() {
// 	fName := flag.String("f", "quiz.csv", "path of csv file")
// 	timer := flag.Int("t", 30, "timer for the quiz")
// 	flag.Parse()
// 	problems, err := problemPuller(*fName)
// 	if err != nil {
// 		handleError(err)
// 	}
// 	correctAns := 0

// 	tObj := time.NewTimer(time.Duration(*timer) * time.Second)

// 	ansc := make(chan string)

// problemloop:
// 	for i, p := range problems {
// 		var answer string
// 		fmt.Printf("%d:%s=", i, p.q)
// 		go func() {
// 			fmt.Scanf("%s", &answer)
// 			ansc <- answer
// 		}()
// 		switch {
// 		case <-tObj.C:
// 			fmt.Println()
// 			break problemloop
// 		case iAns:= <-ansc:
// 			if iAns==p.a{
// 			correctAns++
// 			}
// 			if i==len(problems)-1 {
// 				close(ansc)
// 			}

// 		}
// 	}
// fmt.Printf("Your result is %d out of %d\n",correctAns,len(problems))
// <-ansc
// }

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func problemPuller(fileName string) ([]problem, error) {
	fobj, err := os.Open(fileName)
	if err != nil {
		handleError(err)
		return nil, err
	}
	csvR := csv.NewReader(fobj)
	cLines, err := csvR.ReadAll()
	if err != nil {
		handleError(err)
		return nil, err
	}
	return parseProblem(cLines), nil
}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

func handleError(err error) {
	fmt.Println(err)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	fName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	problems, err := problemPuller(*fName)
	if err != nil {
		handleError(err)
	}

	correctAns := 0
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansc := make(chan string)

problemLoop:
	for i, p := range problems {
		fmt.Printf("%d:%s=", i+1, p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansc <- answer
		}()

		select {
		case <-tObj.C:
			break problemLoop
		case iAns := <-ansc:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansc)
			}
		}
	}
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
}
