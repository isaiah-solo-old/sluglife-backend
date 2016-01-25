package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

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

    _, err = stmtIns.Exec("?", "?", "?")
    if err != nil{
        panic(err.Error())
    }
}

func dbGet(queryType string, query1 string, query2 string){
    db, err := sql.Open("mysql", "user:pw@/database")
    if err != nil{
        panic(err.Error())
    }

    stmtOut,err := db.Prepare("")

    // Expand to handle all types of queries
    if(queryType == "all"){
        // Get all events
        stmtOut, err := db.Prepare("SELECT * FROM event")
        if err != nil{
            panic(err.Error())
        } 

        err = stmtOut.QueryRow("?").Scan("?")
        if err != nil{
            panic(err.Error())
        }       
    }else if(queryType == "date"){
        // Get events within dates
        stmtOut, err := db.Prepare("SELECT * FROM event WHERE date >= ? && date <= ?")
        if err != nil{
            panic(err.Error())
        }    
        
        err = stmtOut.QueryRow("?").Scan("?")
        if err != nil{
            panic(err.Error())
        }
    }

    stmtOut.QueryRow("?").Scan("?") 


}
