package main

import (
	"context"
	db "error-berror/pkg/db/sqlc"
	ex "error-berror/pkg/errors/models"
	handler "error-berror/pkg/errors/handlers"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Connect to the database
	connPool, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5436/db?sslmode=disable")
	if err != nil {
		sysErr, usrErr := handler.HandleException(ex.ErrDBConnectionFailed)
		log.Println("System Error: ", sysErr)
		log.Println("User Error: ", usrErr)
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		sysErr, usrErr := handler.HandleException(ex.ErrDBConnectionFailed)
		log.Println("System Error: ", sysErr)
		log.Println("User Error: ", usrErr)
	}

	// Get a new store
	dbStore := db.NewStore(connPool)

	// Get a new author
	user, err := dbStore.GetAuthor(context.Background(), 1)
	if err != nil {
		sysErr, usrErr := handler.HandleException(err)
		log.Println("System Error: ", sysErr)
		log.Println("User Error: ", usrErr)
	}

	log.Println("User: ", user)
}
