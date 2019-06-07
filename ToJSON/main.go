package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
}

func main()  {
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		var user User
		json.NewDecoder(request.Body).Decode(&user)

		fmt.Fprintf(writer, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname: "Doe",
			Age: 25,
		}

		json.NewEncoder(writer).Encode(peter)
	})

	http.ListenAndServe(":9999",nil);
}


// encoding/json包来编码和解码 JSON文件的.

// curl -s -XPOST -d'{"firstname":"Donald","lastname":"Trump","age":70}' http://localhost:9999/decode Donald Trump is 70 years old!

// curl -s http://localhost:9999/encode