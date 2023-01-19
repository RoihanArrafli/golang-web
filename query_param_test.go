package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T)()  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Roihan", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request)  {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Roihan&last_name=Arrafli", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request)  {
	query := request.URL.Query()
	names := query["name"]
	fmt.Println(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Roihan&name=Arrafli", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _  := io.ReadAll(response.Body)
	fmt.Println(string(body))
}