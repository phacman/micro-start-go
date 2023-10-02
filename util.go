package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func stripFileStatic(dir string) http.Handler {
	pth := fmt.Sprintf("./public/%s/", dir)
	fld := fmt.Sprintf("/%s/", dir)
	fs := http.FileServer(http.Dir(pth))
	return http.StripPrefix(fld, fs)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file := "./public/index.html"
	url := strings.Trim(r.URL.Path, "")

	if url != "/" {
		file = "./public/404.html"
	}

	var tpl = template.Must(template.ParseFiles(file))
	err := tpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["title"] = "Go Micro Start"
	resp["description"] = "Simple microservice with native Go"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
	return
}
