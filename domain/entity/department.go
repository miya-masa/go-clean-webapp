package entity

type Department struct {
	UUID string `json:"uuid" db:"uuid"`
	Name string `json:"name" db:"name"`
}
