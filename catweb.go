package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", CatHandler)
	log.Printf("Listening on port 8080!")
	http.ListenAndServe(":8080", nil)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func CatHandler(w http.ResponseWriter, r *http.Request) {
	//Fetch hostname of container
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Choose random catpicture
	catpic := random(1, 10)

	//Parse index.html template
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}

	items := struct {
		Url      int
		Hostname string
	}{
		Url:      catpic,
		Hostname: name,
	}

	t.Execute(w, items)
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}
