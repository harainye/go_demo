package template

import (
	"log"
	"os"
	"text/template"
)

const templ = ` 
{{range .}}----------------------------------------
Name:   {{.Name}}
Price:  {{.Price | printf "%4s"}}
{{end}}`

var report = template.Must(template.New("report").Parse(templ))

type Book struct {
	Name  string
	Price float64
}

func RenderTemplate() {
	Data := []Book{{"《三国演义》", 19.82}, {"《儒林外史》", 99.09}, {"《史记》", 26.89}}
	if err := report.Execute(os.Stdout, Data); err != nil {
		log.Fatal(err)
	}
}
