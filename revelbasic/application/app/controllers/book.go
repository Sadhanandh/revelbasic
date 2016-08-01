package controllers

import (
    "github.com/sadhanandh/revelbasic"
    "github.com/sadhanandh/revelbasic/application/app/models"
    "github.com/revel/revel"
)

type Book struct {
    *revel.Controller
    revelbasic.MongoController
}

func (c Book) Index() revel.Result {
    return c.Render()
}

func (c Book) View(id string) revel.Result {
    b := models.GetBookById(c.MongoSession, id)
    if b.Id.Hex() != id {
        return c.NotFound("Could not find a book with '%s' as id.", id)
    }

    return c.Render(b)
}
