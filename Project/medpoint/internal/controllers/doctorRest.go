package controllers

import (
	"medpoint/internal/models"

	"github.com/sev-2/raiden"
)

type BooksController struct {
	raiden.ControllerBase
	Http  string `path:"/rest/v1/medpointdoctor" type:"rest"`
	Model models.MedpointDoctor
}
