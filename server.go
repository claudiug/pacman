package main

import (
	"database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	var (
		id    int
		title string
		description string
	)

	// DB STUFF //
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from lessons")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &title, &description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, title, description)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

//	insert, err := db.Prepare("insert into lessons(description, title) values(?, ?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, err := insert.Exec("Dolly", "someother stuff")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(res.RowsAffected())

	defer db.Close()
	//	router := gin.Default()

	//	router.GET("/")
	//	//	LESSONS //
	//	router.GET("/lessons")
	//	router.GET("/lessons/:id")
	//	// USERS //
	//	router.GET("/users")
	//	router.GET("/users/:id")
	//	// LEVELS //
	//	router.GET("/levels")
	//	router.GET("/levels/:id")
	//	// SCORES //
	//	router.GET("/scores")
	//	router.GET("/scores/:id")
	//
	//	router.Run(":8000")
}
