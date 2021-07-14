package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dineshpawar22/booking/cmd/pkg/config"
	"github.com/dineshpawar22/booking/cmd/pkg/handlers"
	"github.com/dineshpawar22/booking/cmd/pkg/render"
)

const portNumber = ":9090"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//Change this to true if you are in PRD
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	sev := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = sev.ListenAndServe()

	log.Fatal(err)
	// _ = http.ListenAndServe(portNumber, nil)
}
