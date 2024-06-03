package main

import (
	"CrudGo/conexao"
	"CrudGo/routes"
)

func main() {
	conexao.InitDB("root@tcp(127.0.0.1:3306)/crud_go")
	routes.IniRoute()
}
