package web

import (
	"net/http"

	"github.com/miya-masa/go-clean-webapp/interface/presenter"
)

type AccountJsonView struct {
	w           http.ResponseWriter
	successCode int
}

func (j *AccountJsonView) ErrorView(err error) {
	j.w.WriteHeader(http.StatusInternalServerError)
}

func (j *AccountJsonView) View(account *presenter.AccountViewModel) {
	if err := jsonView(j.w, j.successCode, account); err != nil {
		j.w.WriteHeader(http.StatusInternalServerError)
	}
}

func (j *AccountJsonView) ViewModels(account []*presenter.AccountViewModel) {
	if err := jsonView(j.w, j.successCode, account); err != nil {
		j.w.WriteHeader(http.StatusInternalServerError)
	}
}

func (j *AccountJsonView) ViewNoBody() {
	j.w.WriteHeader(j.successCode)
}
