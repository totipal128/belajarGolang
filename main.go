package main

import (
	"fmt"
	"golang2/conection"
	"golang2/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println(conection.ConnDB())
	route := routes.GetRoutes()
	route.Run()

}