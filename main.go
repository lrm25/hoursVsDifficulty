package main

import (
    "database/sql"
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    now := time.Now()
    fmt.Printf("Date:  %s\n", now.Format("Jan 2, 2006"))
    createDatabase()
}

// TODO: fix go get/install for this
func createDatabase() {
    db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    dbs, err := db.Query("SHOW DATABASES")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    var dbName string
    for dbs.Next() {
        dbs.Scan(&dbName)
        fmt.Println(dbName)
    }
}
