package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/markstgodard/currency-service/emoji"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query().Get("amount")
	amt, err := strconv.ParseFloat(v, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	code := r.URL.Query().Get("code")

	e, err := emoji.Convert(amt, code)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, e)
}
