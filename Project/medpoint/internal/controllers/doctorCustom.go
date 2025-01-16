package controllers

import (
	"fmt"
	"medpoint/internal/models"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type HelloWordRequest struct { // Add your request data
}

type HelloWordResponse struct {
	Message string `json:"message"`
}

type HelloWordController struct {
	raiden.ControllerBase
	Http    string `path:"/rest/v1/medpointdoctor?select=*&id=eq.1" type:"custom"`
	Payload *HelloWordRequest
	Result  HelloWordResponse
}

func (c *HelloWordController) Get(ctx raiden.Context) error {
	var medpoint_doctor []models.MedpointDoctor

	db.
		NewQuery(ctx).
		From(models.MedpointDoctor{}).
		Select([]string{"id", "name", "isbn"}).
		Eq("isbn", "9786235266008").
		Get()
	// SQL: select id, name, isbn from books where isbn = '9786235266008'
	// URL: /rest/v1/books?select=id,name,isbn&isbn=9786235266008

	fmt.Println(books)
	fmt.Println(books[0].Id)
	fmt.Println(books[0].Name)
	fmt.Println(books[0].Isbn)

	return ctx.SendJson(books)
}
