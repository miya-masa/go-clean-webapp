package entity

import uuid "github.com/satori/go.uuid"

var genUUID = uuid.NewV4

func New(firstName, lastName string) *Account {
	return &Account{
		UUID:      genUUID().String(),
		FirstName: firstName,
		LastName:  lastName,
	}
}
