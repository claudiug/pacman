package main

import (
	"fmt"
	"time"
)

type Question struct {
	Title string
	Answer string
}

type Lesson struct {
	Name string
	Description string
	NumberOfQuestions int
	Questions []Question
}

type Level struct {
	Name string
	Description string
	Lessons []Lesson
}



func main(){
	now := time.Now()
	fmt.Println(now.Date())
	//add level
	level1 := Level{}
	//add lessons
	lessons := make([]Lesson, 3)
	lesson1 := Lesson{}
	lesson1.Name = "First lesson"
	lesson1.Description = "This is a long description"
	//add questions for lesson 1
	questions := make([]Question, 3)
	questions[0]= Question{"A title", "an aswer"}
	questions[1]= Question{"A title1", "an aswer1"}
	questions[2]= Question{"A title2", "an aswer2"}
	lesson1.Questions = questions
	lesson1.NumberOfQuestions = len(lesson1.Questions)
	lessons[0] = lesson1
	level1.Lessons = lessons
	fmt.Println(level1)
}

