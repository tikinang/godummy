package main

import (
	"fmt"
	"net/http"

	"go.elastic.co/apm/v2"
)

func main() {
	fmt.Println(apm.AgentVersion)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "OK")
	})
	_ = http.ListenAndServe(":8080", nil)
}
