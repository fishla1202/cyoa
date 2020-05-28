package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	story, err:= cyoa.ParseJsonFileToStoryType(f)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Println("working on port: 3030")
	log.Fatal(http.ListenAndServe("localhost:3030", h))
}