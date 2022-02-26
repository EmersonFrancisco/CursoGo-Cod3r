package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")
	stmt.Exec("Emerson")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println("ultimo Id inserido foi:", id)

	linhas, err := res.RowsAffected()
	fmt.Println("ultimo comando afetou apenas", linhas, "linha!")
}
