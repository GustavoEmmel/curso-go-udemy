package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values (?, ?)")

	stmt.Exec(2000, "Gustavo")
	stmt.Exec(2001, "Reichelt")
	_, err = stmt.Exec(1, "errou") //chave duplicado

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
