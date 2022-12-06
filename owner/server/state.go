package server

type OwnerState struct {
	Salutation string `json:"salutation" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
}
