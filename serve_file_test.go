package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
	_ "embed"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	}else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServerFileEmbed(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	}else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServerFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}