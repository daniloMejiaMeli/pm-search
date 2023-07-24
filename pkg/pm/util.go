package pm

import "net/http"

func addHeadersToRequest(headers map[string]string) http.Header {
	h := http.Header{}
	for key, value := range headers {
		h.Set(key, value)
	}

	return h
}
