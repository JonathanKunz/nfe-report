package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JonathanKunz/tax-report/src/infrastructure/repository"
	httpHandler "github.com/JonathanKunz/tax-report/src/tax/delivery/http"
	taxUseCase "github.com/JonathanKunz/tax-report/src/tax/usecase"
	"github.com/codegangsta/negroni"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	taxRepo := repository.NewTaxMongoDB()
	taxService := taxUseCase.NewService(taxRepo)

	httpHandler.MakeTaxHandler(r, *n, taxService)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8001),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
