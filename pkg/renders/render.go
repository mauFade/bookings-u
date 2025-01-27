package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedT, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	if err != nil {
		fmt.Println("Error" + err.Error())
	}

	err = parsedT.Execute(w, nil)

	if err != nil {
		fmt.Println("Error" + err.Error())
	}
}
