package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Event struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Parses Request Form
	form_err := r.ParseForm()
	if form_err != nil {
		log.Fatal(form_err)
		return
	}

	// Finds what type of Request 'r' is
	switch r.Method {
		case "GET":
			// Accesses Query String
			name := r.Form.Get("name")
			w.Write([]byte(name))
	 	case "POST":
			var newevent Event
			decoder := json.NewDecoder(r.Body)
			decode_err := decoder.Decode(&newevent)
			if decode_err != nil {
				log.Fatal(decode_err)
			}
			
	 	case "PUT":
	 		w.Write([]byte("put\n"))
		case "DELETE":
			w.Write([]byte("delete\n"))
		default:
	 		w.Write([]byte("default\n"))
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
