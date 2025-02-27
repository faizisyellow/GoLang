package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	UserId    int    `json:"userId"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response header:", resp.Header["Content-Type"])

	var data Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("%+v\n", data)
}
