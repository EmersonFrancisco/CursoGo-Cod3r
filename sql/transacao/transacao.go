package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // dessa forma ele inicia uma transação, para poder analisar antes de jogar para o banco de dados
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	// exempolo de roolback abaixo

	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// dessa forma como deu um erro, nem a Bia e nem o Carlos serão inseridos no BD,
	// pois como ocorreu um erro, ele cancelou todatransação!!!!

	tx.Commit() // execulta todo a transação registrada
}
