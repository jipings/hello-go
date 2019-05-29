package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main()  {
	tmpl, err := template.ParseFiles("layout.html")
	panic(err);
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todo list",
			Todos: []Todo{
				{ Title: "Task 1", Done: false },
				{ Title: "Task 2", Done: true },
				{ Title: "Task 3", Done: true },
			},
		}
		tmpl.Execute(writer, data)
	})

	http.ListenAndServe(":9999", nil);
}