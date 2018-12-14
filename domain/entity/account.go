package entity

type Account struct {
	UUID      string `json:"uuid" db:"uuid"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}
