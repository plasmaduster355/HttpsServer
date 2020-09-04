package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Files struct {
	CrtFileName string `json:"CrtFileName"`
	KeyFileName string `json:"KeyFileName"`
}

func main() {
	var files Files
	b, _ := ioutil.ReadFile("data.json")
	json.Unmarshal(b, &files)
	fs := http.FileServer(http.Dir("./site"))
	http.ListenAndServeTLS(":8082", files.CrtFileName, files.KeyFileName, fs)
}
