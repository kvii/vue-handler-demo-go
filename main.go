package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/kvii/handler"
)

//go:embed dist
var dist embed.FS

func main() {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler.VueServer(http.FS(sub)))
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	log.Println("Serve at http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
