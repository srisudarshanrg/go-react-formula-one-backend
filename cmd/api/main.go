package main

import (
	"fmt"
	"net/http"
)

const port = 10000

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello from formula one new backend")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
