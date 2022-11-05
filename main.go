package main

import (
    "elstore/configs"
	 "elstore/routes" 
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // run database
    configs.ConnectDB()

	// reoutes
	routes.UserRoute(e)

    e.Logger.Fatal(e.Start(":6000"))
}