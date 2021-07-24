package main

import "net/http"

func main() {

	http.HandleFunc("/",Index)

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request)  {

}