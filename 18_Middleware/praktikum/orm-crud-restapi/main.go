package main

import (
	"orm-crud-restapi/config"
	"orm-crud-restapi/routes"
)

func main() {
	config.Init()
	e := routes.New()
	
	e.Logger.Fatal(e.Start(":8000"))
	
}
