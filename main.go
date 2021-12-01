package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kazan-emercom-api/parser"
	"log"
	"net/http"
	"os"
	"strings"
)

func getHandler(fileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.Open("data/" + fileName + ".txt")
		defer file.Close()

		objectsArray := parser.Parser(file)

		b, _ := json.Marshal(objectsArray)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		fmt.Fprintf(w, string(b))
	}
}

type NewKazanObject struct {
	File        string
	Coordinates [2]string
	Name        string
	Extra       string
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var newObject NewKazanObject

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &newObject)

	fmt.Println(newObject)

	file, _ := os.OpenFile("data/"+newObject.File+".txt", os.O_APPEND|os.O_RDWR, 0600)
	defer file.Close()

	file.WriteString(newObject.Coordinates[0] + ", " +
		newObject.Coordinates[1] + ", " +
		newObject.Name + ", " +
		newObject.Extra + "\n")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "OK")
}

func main() {
	files, _ := ioutil.ReadDir("data")

	for _, file := range files {
		if file.Name() != ".DS_Store" {
			fileName := strings.Split(file.Name(), ".")
			http.HandleFunc("/"+fileName[0], getHandler(fileName[0]))
		}
	}

	http.HandleFunc("/create", createHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
