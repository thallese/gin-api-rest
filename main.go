package main

import (
	"github.com/creycrey/api-go-gin/database"
	"github.com/creycrey/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
