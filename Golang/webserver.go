package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Files struct {
	CrtFileName string `json:"CrtFileName"`
	KeyFileName string `json:"KeyFileName"`
	Port        int    `json:"Port"`
}

func main() {
	var files Files
	b, _ := ioutil.ReadFile("data.json")
	json.Unmarshal(b, &files)
	fs := http.FileServer(http.Dir("./site"))
	http.ListenAndServeTLS(":"+strconv.Itoa(files.Port), files.CrtFileName, files.KeyFileName, fs)
}
