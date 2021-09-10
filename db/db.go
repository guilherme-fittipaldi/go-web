package db

import (
	"database/sql" 	
	
_ "github.com/lib/pq"
)

func ConectaComBandoDeDados() *sql.DB {
	conexao := "user=postgres dbname=loja password=new_password host=localhost sslmode=disable"
	
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}