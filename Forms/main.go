package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func main()  {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
			return
		}

		details := ContactDetails{
			Email: request.FormValue("email"),
			Subject: request.FormValue("subject"),
			Message: request.FormValue("message"),
		}

		_ = details

		tmpl.Execute(writer, struct {
			Success bool
		}{true})
	})

	http.ListenAndServe("9999",nil)
}
