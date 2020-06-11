package main

import (
	"net/http"
)

func main() {
	DownloadWasmExec()
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)
	println("Listening on port 5000...")
	http.ListenAndServe(":5000", nil)
}
