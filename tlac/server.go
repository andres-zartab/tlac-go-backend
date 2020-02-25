package tlac

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type credentials struct {
	host, user, password, dbname string
	port                         int
}

type server struct {
	db     *sql.DB
	router *mux.Router
	//email  EmailSender
	logger *log.Logger
}

func NewDB() (*sql.DB, error) {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func newServer() *server {
	s := &server{}
	//s.routes()
	return s
}
