package main

import (
	"net/http"
	"time"
	"github.com/altuntasfatih/simple-golang-restful-with-docker-and-postgres/muxes"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/altuntasfatih/simple-golang-restful-with-docker-and-postgres/config"
)

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.PORT)
	db, err := sql.Open("postgres", dbinfo)
	log.Printf("Eror check",err )
	checkErr(err)
	log.Printf("Postgres started at %s PORT", config.PORT)
	defer db.Close()

	s := &http.Server{
		Addr:           ":3000",
		Handler:        muxes.SERVE(db),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func checkErr(err error) {
	if err != nil {
		log.Printf("Eror check in main checkeRR",err )
		panic(err)
	}
}