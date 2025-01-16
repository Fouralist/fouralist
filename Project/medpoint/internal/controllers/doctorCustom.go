package controllers

import (
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
		// Select([]string{"id", "name_doctor"}).
		Eq("id", 1).
		Single()
	// SQL: select id, name, isbn from books where isbn = '9786235266008'
	// URL: /rest/v1/books?select=id,name_doctor&id=eq.1

	// fmt.Println(medpoint_doctor)
	// fmt.Println(medpoint_doctor[0].Id)
	// fmt.Println(medpoint_doctor[0].Id_Occupation)
	// fmt.Println(medpoint_doctor[0].Occupation_Doctor)
	// fmt.Println(medpoint_doctor[0].Name_Doctor)
	// fmt.Println(medpoint_doctor[0].Contact_Doctor)

	return ctx.SendJson(medpoint_doctor)
}
