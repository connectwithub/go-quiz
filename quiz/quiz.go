package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

//StartQuiz starts the quiz and return the score
func StartQuiz(fileName string, timeLimit int) int {
	score := 0
	file, err := os.Open(path.Join("quiz", fileName))
	if err != nil {
		log.Fatalln("File '" + fileName + "' Not Found")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to read file")
	}
	problems, answers := parseProblems(lines)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for i, p := range problems {
		answerCh := make(chan string)
		go func(answerCh chan<- string) {
			fmt.Printf("Question: %s \n", p)
			reader := bufio.NewReader(os.Stdin)
			ans, _ := reader.ReadString('\n')
			ans = strings.Trim(ans, " \n")
			answerCh <- ans
		}(answerCh)
		select {
		case <-timer.C:
			fmt.Printf("Timer ended\n")
			return score
		case ans := <-answerCh:
			if ans == answers[i] {
				score += 10
			} else {
				score -= 5
			}
		}
	}
	return score
}

func parseProblems(file [][]string) ([]string, []string) {
	problems := []string{}
	solutions := []string{}
	for _, row := range file {
		problems = append(problems, row[0])
		solutions = append(solutions, row[1])
	}
	return problems, solutions
}
