package template

import (
	"net/http"
	"text/template"
)

func tHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Hello World!")
}

func RenderHTemplate() {
	http.HandleFunc("/", tHandler)
	http.ListenAndServe(":8089", nil)
}
