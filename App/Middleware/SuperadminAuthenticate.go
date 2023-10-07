package Middleware

import (
	"github.com/globalxtreme/gobaseconf/data"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

func SuperadminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employee := data.Employee
		if !employee.Superadmin {
			error.ErrXtremeUnauthenticated("Your access must be superadmin")
		}

		next.ServeHTTP(w, r)
	})
}
