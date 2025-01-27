package main

import (
	"fmt"
	"net/http"

	"github.com/mauFade/bookings-u/pkg/handlers"
)

const PORT = ":9091"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
