package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
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

func dbInsert(){
    db, err := sql.Open("mysql", "user:pw@/database")
    if err != nil{
        panic(err.Error())
    }
    defer db.Close()

    stmtIns, err := db.Prepare("INSERT INTO event VALUES(?,?,?)")
    if err != nil{
        panic(err.Error())
    }
    defer stmtIns.Close()

    _, err = stmtIns.Exec("?", "?")
    if err != nil{
        panic(err.Error())
    }
}

func dbGet(){
    db, err := sql.Open("mysql", "user:pw@/database")
    if err != nil{
        panic(err.Error())
    }

    // Expand to handle all types of queries
    stmtOut, err := db.Prepare("SELECT * FROM event WHERE date = ?")
    if err != nil{
        panic(err.Error())
    }

    var event Event 
    err = stmtOut.QueryRow("?").Scan(&event)
    if err != nil{
        panic(err.Error())
    }
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
