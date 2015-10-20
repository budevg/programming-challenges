package main

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage %s <port>\n", os.Args[0])
		os.Exit(1)
	}

	addr := "localhost:" + os.Args[1]

	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadFiles)
	mux.HandleFunc("/mosaic", createMosaic)

	server := http.Server{
		Addr : addr,
		Handler : mux,
	}

	fmt.Printf("Starting listening on %s\n", addr)
	server.ListenAndServe()

}


func uploadFiles(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createMosaic(w http.ResponseWriter, req *http.Request) {
}
