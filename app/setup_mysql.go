package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func setupMysql(){
	// data source name
	var dsn string = globalConf["dbUser"] + ":" + globalConf["dbPassword"] + "@(" + globalConf["dbHost"] + ":" + globalConf["dbPort"] + ")/" + globalConf["dbDB"]
	var err error
	log.Println("Used URI:",dsn)

	dbInfo.db, err = sql.Open("mysql", dsn)

	if err != nil{
		log.Fatal(err)
	}

	if err := dbInfo.db.Ping(); err != nil {
		// Waiting that the Database is fully up
		log.Fatal("Connection is not established")
	}

	setupTables()

}


func setupTables(){
	stmt, err := dbInfo.db.Prepare("CREATE TABLE memes(id INTEGER AUTO_INCREMENT, link VARCHAR(512), post_time DATE, path VARCHAR(45), creator INTEGER, PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}





}
