package main

import (
	"courses/internal/app"
	"fmt"
)

// @title Courses - Course management system
// @version 1.0
// @description This is a simple course management system

// @host coursestj.herokuapp.com
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("Starting the programm ...")

	confpath := "config/conf.json"

	err := app.Run(confpath)
	if err != nil {
		fmt.Println(err)
	}
}
