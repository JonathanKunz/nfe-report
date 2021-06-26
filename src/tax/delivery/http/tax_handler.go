package http

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/JonathanKunz/tax-report/src/domain"
	"github.com/JonathanKunz/tax-report/src/tax/usecase"
)

func fetchTax(service usecase.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading book"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorMessage))
		return
	})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.InternalServerError:
		return http.StatusInternalServerError
	case domain.NotFound:
		return http.StatusNotFound
	case domain.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func MakeTaxHandler(r *mux.Router, n negroni.Negroni, service usecase.UseCase) {
	r.Handle("/tax-report", n.With(
		negroni.Wrap(fetchTax(service)),
	)).Methods("GET", "OPTIONS").Name("fetchTax")
}
