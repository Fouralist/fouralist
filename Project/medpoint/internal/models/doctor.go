package models

import (
	"github.com/sev-2/raiden"
)

type MedpointDoctor struct {
	raiden.ModelBase
	Id                string `json:"id"`
	Id_Occupation     string `json:"id_occupation"`
	Occupation_Doctor string `json:"Occupation"`
	Name_Doctor       string `json:"Name Doctor"`
	Contact_Doctor    string `json:"Contact Doctor"`

	Metadata string `json:"-" schema:"public"`

	Acl string `json:"-" read:"" write:""`
}
