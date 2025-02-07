package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mauFade/bookings-u/pkg/config"
	"github.com/mauFade/bookings-u/pkg/handlers"
	"github.com/mauFade/bookings-u/pkg/renders"
)

const PORT = ":9091"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := renders.CreateTemplate()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	fmt.Printf("Staring application on port %s", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
