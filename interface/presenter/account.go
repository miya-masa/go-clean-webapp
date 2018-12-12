package presenter

import "github.com/miya-masa/go-clean-webapp/domain/entity"

type AccountViewModel struct {
	UUID           string `json:"uuid"`
	DepartmentUUID string `json:"department_uuid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

type AccountPresenter struct {
}

func (a *AccountPresenter) ToViewModel(ac *entity.Account) *AccountViewModel {
	return &AccountViewModel{
		UUID:           ac.UUID,
		DepartmentUUID: ac.Department.UUID,
		FirstName:      ac.FirstName,
		LastName:       ac.LastName,
	}
}
