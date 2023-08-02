package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/accounting-api/handlers"
	"gitlab.sudovi.me/erp/accounting-api/middleware"

	"github.com/oykos-development-hub/celeritas"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/accounting-api"

	myHandlers := &handlers.Handlers{}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
