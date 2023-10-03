package Middleware

import (
	"github.com/globalxtreme/gobaseconf/data"
	"github.com/globalxtreme/gobaseconf/response/error"
	"log"
	"net/http"
)

func SuperadminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employee := data.Employee
		if !employee.Superadmin {
			log.Println(employee.Superadmin)
			error.ErrUnauthenticated("Your access must be superadmin")
		}

		next.ServeHTTP(w, r)
	})
}
