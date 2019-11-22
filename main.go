package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nS := r.URL.Query().Get("n")
		if nS == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		n, err := strconv.Atoi(nS)
		if err != nil || n < 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		res := Fact(uint(n))
		fmt.Fprintf(w, "Fact(%d) = %d\n", n, res)
	})

	http.ListenAndServe(":8080", nil)
}
