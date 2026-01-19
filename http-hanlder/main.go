package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != "POST" {
		// http.Error(response, "Method is not supported", http.StatusMethodNotAllowed)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusMethodNotAllowed)
		messageError := map[string]string{
			"message": "Method is not supported",
		}
		json.NewEncoder(response).Encode(messageError)
		return
	}

	var p Person
	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	responseData := map[string]string{
		"message": "Hello " + p.Name,
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(responseData)

	// fmt.Fprint(response, "Hello Handler")
	// fmt.Println(p.Name)
	// log.Print(p.Name)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
