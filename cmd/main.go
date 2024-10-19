package main

import (
	"database/sql"
	"log"

	"github.com/KairongChen-l/Gose/cmd/api"
	"github.com/KairongChen-l/Gose/config"
	"github.com/KairongChen-l/Gose/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db,err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassWord,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err!=nil{
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080",db)
	if err := server.Run();err!=nil{
		log.Fatal(err)
	}
}
func initStorage(db *sql.DB){
	err:=db.Ping()
	if err!= nil{
		log.Fatal(err)
	}
	log.Println("DB: Seccessfully connected!")
}