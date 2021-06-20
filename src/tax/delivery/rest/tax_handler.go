package rest

import (
	"net/http"

	"github.com/JonathanKunz/tax-report/src/domain/"
)

func FetchTax() string {
	return "test"
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
