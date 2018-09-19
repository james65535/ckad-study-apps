package main

import (
	"io/ioutil"
	"net/http"
)

func serveContent(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("/data/dat")
	if err != nil{
		panic(err)
	}
	w.Write([]byte(f))
}
func main() {
	http.HandleFunc("/", serveContent)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
