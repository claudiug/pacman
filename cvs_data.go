package main

import (
	"encoding/csv"
	"log"
	"os"
	"io"
	"fmt"
)

func main() {

	file, err := os.Open("filelocation")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			log.Fatal(err)
			return
		}else if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

}
