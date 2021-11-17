package controller

import (
	"github.com/labstack/echo"
	"github.com/yhosec/notes-api-go/config"
	"github.com/yhosec/notes-api-go/dto"
	"net/http"
)

func NotesIndex(context echo.Context) error {
	db := config.ConnectMongo()
	db.Collection("notes")
	return context.HTML(200, "Hello")
}

func NotesCreate(context echo.Context) error {
	createNotesRequest := new(dto.CreateNotesRequest)
	if err := context.Bind(createNotesRequest); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, createNotesRequest)
}
func NotesDelete(context echo.Context) error {
	return context.HTML(200, "Hello")
}
