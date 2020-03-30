package main

import (
	_ "feita/databases"
	"feita/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8095")
}
