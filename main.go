package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const PORT = ":9091"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedT, err := template.ParseFiles("./templates" + tmpl)

	if err != nil {
		fmt.Println("Error" + err.Error())
	}

	err = parsedT.Execute(w, nil)

	if err != nil {
		fmt.Println("Error" + err.Error())
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
