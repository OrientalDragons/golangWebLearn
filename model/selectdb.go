package model

import (
	"database/sql"
	"fmt"
)

func ConnDb() *sql.DB {
	return db
}

//TestDb  xx
func TestDb(db *sql.DB) string {
	rows, err := db.Query("select * from user")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var id int64
	var name, psw string
	rows.Next()
	rows.Scan(&id, &name, &psw)
	fmt.Println(id, name, psw)

	return name
}
