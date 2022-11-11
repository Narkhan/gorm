package main

import (
	"final/models"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "hospitals")
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	checkError(err)
	fmt.Println("Successfully connected!")
	db.AutoMigrate(&models.User{}, &models.Pharmacy{}, &models.Drug{})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
