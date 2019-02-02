package presenter

import "github.com/miya-masa/go-clean-webapp/domain/entity"

type AccountViewModel struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AccountPresenter struct {
}

func (a *AccountPresenter) ToViewModel(ac *entity.Account) *AccountViewModel {
	return &AccountViewModel{
		UUID:      ac.UUID,
		FirstName: ac.FirstName,
		LastName:  ac.LastName,
	}
}

func (a *AccountPresenter) ToViewModels(ac []*entity.Account) []*AccountViewModel {
	res := make([]*AccountViewModel, 0, len(ac))
	for _, model := range ac {
		res = append(res, a.ToViewModel(model))
	}
	return res
}
