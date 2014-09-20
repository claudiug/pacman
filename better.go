package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	type Level struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Points      int    `json:"points"`
	}

	type Question struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Response  string `json:"response"`
		Points    int    `json:"value"`
		Lesson_Id int    `json:"lesson_id"`
	}

	type Lesson struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Points      int    `json:"points"`
		Status      int    `json:"status"`
		Level_Id    int    `json:"level_id"`
	}

	db := sqlx.MustConnect("mysql", "root@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer db.Close()

	router := gin.Default()

	//LIST ALL LESSONS

	router.GET("/lessons", func(c *gin.Context) {
		lessons := []Lesson{}
		rows, _ := db.Queryx("SELECT * from lessons")
		for rows.Next() {
			lesson := Lesson{}
			rows.StructScan(&lesson)
			lessons = append(lessons, lesson)
		}
		rows.Close()
		c.JSON(200, lessons)
	})

	//RETURN ALL LEVELS

	router.GET("/levels", func(c *gin.Context) {
		levels := []Level{}
		rows, _ := db.Queryx("SELECT * from levels")
		for rows.Next() {
			level := Level{}
			rows.StructScan(&level)
			levels = append(levels, level)
		}

		rows.Close()

		c.JSON(200, levels)
	})

	//RETURN ALL QUESTIONS

	router.GET("/questions", func(c *gin.Context) {
		questions := []Question{}
		rows, _ := db.Queryx("SELECT * from questions")
		for rows.Next() {
			question := Question{}
			rows.StructScan(&question)
			questions = append(questions, question)
		}
		rows.Close()

		c.JSON(200, questions)
	})

	//Return JUST ONE LEVEL

	router.GET("/levels/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		fmt.Println(id)
		level := Level{}
		db.Get(&level, "select * from levels where id=?", id)
		c.JSON(200, level)
	})

	//Return JUST ONE QUESTION

	router.GET("/questions/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		fmt.Println(id)
		question := Question{}
		db.Get(&question, "select * from questions where id=?", id)
		fmt.Println(question)
		c.JSON(200, question)
	})

	//Return JUST ONE LESSON

	router.GET("/lessons/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		fmt.Println(id)
		lesson := Lesson{}
		db.Get(&lesson, "select * from lessons where id=?", id)
		fmt.Println(lesson)
		c.JSON(200, lesson)
	})

	router.PUT("/lessons/status/:value/:id", func(c *gin.Context) {
		value := c.Params.ByName("value")
		id := c.Params.ByName("id")
		query := "update lessons set level_id=? where id=?"
		result,_ := db.MustExec(query, value, id).RowsAffected()
		fmt.Println(value)
		c.JSON(200,result)
	})

	router.Run(":3001")
}
