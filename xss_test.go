package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var myTemplate = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

// func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
// 	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
// 		"Title": "Go-lang auto escape",
// 		"Body": "<p>Ini Adalah Body</p>",
// 	})
// }

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request)  {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "GO-Lang Auto Escape",
		"Body": "<p>Ini Adalah Body<script>alert('Anda di hack')</script></p>",
	})
}

func TestAutoEscape(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request)  {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "GO-Lang Auto Escape",
		"Body": template.HTML("<p>Ini Adalah Body</p>"),
	})
}

func TestAutoEscapeDisabled(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabled(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request)  {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "GO-Lang Auto Escape",
		"Body": template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}