package main

import (
	"fmt"
	"icl-broker/pkg/adapter/rest"
	"log"

	"github.com/labstack/echo/v4"
)

const webPort = "80"

type Config struct{}

func main() {
	// app := Config{}

	log.Printf("Starting broker on port %s\n", webPort)

	e := echo.New()
	e = rest.NewRouter(e)

	fmt.Println("Server listen at http://localhost:80")
	if err := e.Start(":80"); err != nil {
		log.Fatalln(err)
	}

	// //define http server
	// srv := &http.Server{
	// 	Addr: fmt.Sprintf(":%s", webPort),
	// 	Handler: app.routes(),
	// }

	// // start web server
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Panic(err)
	// }
}
