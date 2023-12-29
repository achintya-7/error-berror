package main

import (
	err "error-berror/pkg/errors"
	"errors"
	"log"
)

func main() {
	log.Println(err.HandleException(err.ErrDBQueryFailed))
	log.Println(err.HandleException(err.ErrGcpPubsubFailed))
	log.Println(err.HandleException(errors.New("some error")))
}
