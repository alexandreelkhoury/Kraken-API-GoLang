package main

import (
	"database/sql"
	"log"
)

func DataBase() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=mypassword dbname=krakendb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert data into the "status" table
	status := getStatus()
	_, err = db.Exec("INSERT INTO status (status, timestamp) VALUES ($1, $2)",
		status.Result.Status, status.Result.Timestamp)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the "time" table
	time := getTime()
	_, err = db.Exec("INSERT INTO time (unixtime, rfc1123) VALUES ($1, $2)",
		time.Result.Unixtime, time.Result.Rfc1123)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the "assets" table
	/*assets := getAssets()
	_, err = db.Exec("INSERT INTO assets(result) VALUES($1)",
		assets.Result)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the "assetPairs" table
	assetPairs := getAssetPairs()
	_, err = db.Exec("INSERT INTO assetPairs(result) VALUES($1)",
		assetPairs.Result)
	if err != nil {
		log.Fatal(err)
	}
	createXMLFile(myAssets)*/
}
