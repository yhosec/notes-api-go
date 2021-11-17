package main

import (
	"github.com/labstack/echo"
	"github.com/yhosec/notes-api-go/controller"
)

func main()  {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return controller.NotesIndex(context)
	})
	e.POST("/", controller.NotesCreate)
	e.Start(":1322")
}
