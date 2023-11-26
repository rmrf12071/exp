package lib

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	// handle: simple file server
	wd, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(wd+"/public")))

	// start http server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
