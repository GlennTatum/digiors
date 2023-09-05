package models

type AnonUser struct {
	Model
	Name              string `json:"name"`
	RegistrationRefer string
}
