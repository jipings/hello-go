package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Println(vars);
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(writer, "You've requested the book: %s on page %s \n", title, page)
	})

	http.ListenAndServe(":9999",r);
}

/*
https://books.studygolang.com/gowebexamples/routes-using-gorilla-mux/
*/