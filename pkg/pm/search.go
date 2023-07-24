package pm

import (
	"fmt"
	"net/http"
)

// HeadersValidation function to validate headers of the request
func HeadersValidation(r *http.Request) {
	isInternal := isInternalTraffic(r)
	// TODO Imprimir todas las variables que se estan usando en el metodo
	fmt.Printf("isInternal: %v\n", isInternal)
	fmt.Printf("headersValidation function\n")
}

func isInternalTraffic(r *http.Request) bool {
	return r.Header.Get("X-Public") == "true"
}
