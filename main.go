package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type GameState struct {
	Name string
	Points int
	Questions []Question
}

type Question struct {
	Text string
	Option []string
	Answer int
}

func (g *GameState) Init() {
	fmt.Println("Welcome to the quiz game!")
	fmt.Println("Please enter your name: ")
	reader := bufio.NewReader((os.Stdin))

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Error reading name: " + err.Error())
	}

	g.Name = name

	fmt.Printf("let's the game! %s", g.Name)

	}

	func (g *GameState) ProcessCSV() {
		f, err := os.Open("question_go.csv")

		if err != nil {
			panic("Error opening file")
		}
		
		defer f.Close()

		reader := csv.NewReader(f)

		records, err := reader.ReadAll()

		if err != nil {
			panic("Error reading file")
		}

		var allQuestions []Question

		for index, record := range records {
			if index > 0 {
				correctAnswer, _ := toInt(record[5])
				Question := Question{
					Text:  record[0],
					Option: record[1:5],
					Answer: correctAnswer,
				}

				allQuestions = append(allQuestions, Question)
			}
		}

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		rng.Shuffle(len(allQuestions), func(i, j int) {
			allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
		})

		if len(allQuestions) < 10 {
			panic("Warning: not enough question in the CSV file. Using all available questions.")
		} else {
			g.Questions = allQuestions[:10]
		}

	}

	func (g *GameState) Run() {
			timeLimit := 5 * time.Minute

			timeout := time.After(timeLimit)

			for i, question := range g.Questions {
					fmt.Println("")
					fmt.Printf("\033[32m%d. %s\033[0m\n", i+1, question.Text)

					fmt.Println("------------------------------")
					for j, option := range question.Option {
							fmt.Printf("[%d] %s \n", j+1, option)
					}

					fmt.Println("------------------------------")
					fmt.Println("Please enter your answer: ")

					answerChan := make(chan int)
					go func() {
							var answer int
							var err error
							for {
									reader := bufio.NewReader(os.Stdin)
									read, _ := reader.ReadString('\n')
									answer, err = toInt(read[:len(read)-1])
									if err != nil {
											fmt.Println(err.Error())
											continue
									}
									answerChan <- answer
									break
							}
					}()

					select {
					case <-timeout:
							fmt.Println("\nTime's up! The game is over.")
							fmt.Printf("You scored %d points.\n", g.Points)
							return
					case answer := <-answerChan:
							if answer == question.Answer {
									fmt.Println("------------------------------")
									fmt.Println("\033[32m Congratulations!! Your answer is correct! \033[0m")
									fmt.Println("------------------------------")
									g.Points += 10
							} else {
									fmt.Println("------------------------------")
									fmt.Println("\033[31m Ops... your answer is wrong \033[0m")
									fmt.Println("------------------------------")
							}
					}

					time.Sleep(2 * time.Second)

					fmt.Print("\033[H\033[2J")
			}

			fmt.Printf("End game, you craft %d points.\n", g.Points)
	}

	func (g *GameState) ShowRules() {
    fmt.Println("------------------------------")
    fmt.Println("Welcome to the Quiz Game!")
    fmt.Println("Rules:")
    fmt.Println("1. You will be asked 10 random questions.")
    fmt.Println("2. Each correct answer gives you 10 points.")
    fmt.Println("3. You have a limited time the 5 minutes to answer all questions.")
    fmt.Println("4. Enter the number corresponding to your answer.")
    fmt.Println("5. The game ends when time runs out or all questions are answered.")
    fmt.Println("------------------------------")
}

func main() {
	game := GameState{}

	game.ShowRules()

	go game.ProcessCSV()
	
	game.Init()

	game.Run()
}

func toInt(str string) (int, error) {
	i, err := strconv.Atoi(str)

	if err != nil {
		return 0, errors.New("This character is not permitted, please insert a number.")
	}

	return i, nil
}
