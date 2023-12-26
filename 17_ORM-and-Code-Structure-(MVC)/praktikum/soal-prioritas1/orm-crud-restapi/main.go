package main

import (
	"orm-crud-restapi/config"
	"orm-crud-restapi/route"
)

func main() {
	config.Init()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
	
}
