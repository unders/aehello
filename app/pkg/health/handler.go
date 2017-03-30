package health

import (
	"fmt"
	"net/http"
)

func Handler() (string, func(http.ResponseWriter, *http.Request)) {
	return "/_ah/health", func(w http.ResponseWriter, q *http.Request) {
		fmt.Fprint(w, "OK")
	}
}
