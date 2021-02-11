package main

import (
	"flag"
	"fmt"

	"github.com/connectwithub/grophercises/quiz"
)

func main() {
	fileName := flag.String("file", "problems.csv", "Provide the filename to be used (default problems.csv)")
	timeLimit := flag.Int("limit", 30, "Time Limit in Seconds (default 30)")
	flag.Parse()
	score := quiz.StartQuiz(*fileName, *timeLimit)
	fmt.Printf("Score is: %v \n", score)
}
