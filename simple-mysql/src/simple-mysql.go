package main

import (
   "fmt"	
   "log"
   "database/sql"
   _ "github.com/go-sql-driver/mysql" 
)

func main() {
    db, err := sql.Open("mysql","testing:testing@tcp(localhost:3306)/testing")
    if(err != nil) {
        log.Panic(err.Error())
	}
	
	defer db.Close()

	rows, err := db.Query("select fld_1 from test1")
    if(err != nil) {
        log.Panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var ticket_num string

		rows.Scan(&ticket_num)
		fmt.Println(ticket_num)
	}

}