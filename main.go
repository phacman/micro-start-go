package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pin := os.Getenv("PORT_INTERNAL")
	if pin == "" {
		pin = "5100"
	}
	pex := os.Getenv("PORT_EXTERNAL")
	srv := fmt.Sprintf("Server Ports: %s:%s", pex, pin)
	log.Println(srv)
	mux := http.NewServeMux()
	mux.Handle("/assets/", stripFileStatic("assets"))
	mux.Handle("/download/", stripFileStatic("download"))
	mux.HandleFunc("/api", handleApi)
	mux.HandleFunc("/", handleIndex)
	wrappedMux := middlewareDownload(mux)
	log.Fatal(http.ListenAndServe(":"+pin, wrappedMux))
}

func middlewareDownload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/download/") {
			fn := filepath.Base(r.URL.Path)
			fp := filepath.Join("./public/download/", fn)
			log.Println("Downloaded File: ", fp)
			w.Header().Set("Content-Disposition", "attachment; filename="+fn)
		}
		next.ServeHTTP(w, r)
	})
}
