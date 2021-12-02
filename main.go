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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getHandler(fileName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.Open("data/" + fileName + ".txt")
		defer file.Close()

		objectsArray := parser.Parser(file)

		b, _ := json.Marshal(objectsArray)

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

	fmt.Fprintf(w, "OK")
}

func main() {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	files, _ := ioutil.ReadDir("data")

	for _, file := range files {
		if file.Name() != ".DS_Store" {
			fileName := strings.Split(file.Name(), ".")
			router.HandleFunc("/"+fileName[0], getHandler(fileName[0]))
		}
	}

	router.HandleFunc("/create", createHandler)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
