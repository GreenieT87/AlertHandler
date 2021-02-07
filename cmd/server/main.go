package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/GreenieT87/AlertHandler/internal/transport/http"
)

//App -
type App struct{}

// Run - sets up Things
func (app *App) Run() error {
	fmt.Println("setting things up")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
	}
	return nil
}

func main() {
	fmt.Println("AlertHandler")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up")
		fmt.Println(err)
	}
}
