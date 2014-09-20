package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "log"
)

func main34() {

	type Lesson struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// DB STUFF //
	//
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("select * from lessons")
		if err != nil {
			log.Fatal(err)
		}

		results := []Lesson{}
		for rows.Next() {
			result := Lesson{}
			err := rows.Scan(&result.Id, &result.Title, &result.Description)
			results = append(results, result)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.JSON(200, results)
		defer db.Close()
		fmt.Println(results)
	})
	router.Run(":3001")
}
