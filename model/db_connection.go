package model

import (
	"database/sql"
	"fmt"
	Dimo "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP01/struct"
	"time"
	// Import for postgres
	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectToDB Set up connection to the postgres DB
// Will panic on error
func ConnectToDB(host string, dbname string, user string, password string, port int64) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	tempDB, err := sql.Open("postgres", psqlInfo)

	// Open up our database connection.
	// set up a database on local machine using phpmyadmin.
	// The database is called gomvc

	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	err = tempDB.Ping()

	numberOfTest := 0

	for err != nil && numberOfTest < 5 {
		fmt.Println(err)
		fmt.Println("Connection to DB did not succeed, new try")

		time.Sleep(5 * time.Second)
		tempDB, err = sql.Open("postgres", psqlInfo)
		err = tempDB.Ping()

		numberOfTest++
	}

	if err != nil {
		fmt.Println("Database initialisation error")
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	// defer the close till after the main function has finished
	// executing
	db = tempDB
}

func connectDb()  {
	ConnectToDB("db-postgresql-nyc3-10435-do-user-9402464-0.b.db.ondigitalocean.com", "defaultdb", "doadmin", "S8UO9iNtyaYIaHzL", 25060)
}

func GetAllPlayers() []Dimo.Player  {
	connectDb()
	rows, err := db.Query("SELECT * FROM player")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pls []Dimo.Player

	for rows.Next() {
		var pl Dimo.Player

		err := rows.Scan(&pl.PlayerId, &pl.Name, &pl.Avatar, &pl.DiscordId)
		if err != nil {
			panic(err)
		}

		pls = append(pls, pl)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return pls
}

func GetAllRounds() []Dimo.Round  {
	connectDb()
	rows, err := db.Query("SELECT * FROM round")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var rds []Dimo.Round

	for rows.Next() {
		var rd Dimo.Round

		err := rows.Scan(&rd.GameId, &rd.PlayerId, &rd.Reason, &rd.Word, &rd.SubmittedAt)
		if err != nil {
			panic(err)
		}

		rds = append(rds, rd)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return rds
}

// CloseDbConnection will end dialogue with the DB
// Recommanded to use at program's end
func CloseDbConnection(db *sql.DB) {
	defer db.Close()
	fmt.Println("DB is closed")
}
