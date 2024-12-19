package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var k int

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Fprintf(w, "%d", k)
	} else if r.Method == "POST" {
		delta := r.Form.Get("count")
		idelta, err := strconv.Atoi(delta)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "это не число")
			return
		}
		k += idelta
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Метод не поддерживается")
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}
}
