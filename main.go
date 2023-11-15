package main

import (
	"log"

	"mangalist/controllers"
	"mangalist/database"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Database connection
	db, err := sqlx.Connect("postgres", "")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(database.Schema)

	e.GET("/", controllers.Ping)
	e.Logger.Fatal(e.Start(":8000"))
}
