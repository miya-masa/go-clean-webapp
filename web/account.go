package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/miya-masa/go-clean-webapp/interface/presenter"
	"github.com/miya-masa/go-clean-webapp/usecase"
)

type AccountHandler struct {
	Usecase   usecase.AccountInputPort
	Presenter *presenter.AccountPresenter
}

func NewAccountHandler(uc usecase.AccountInputPort, pr *presenter.AccountPresenter) *AccountHandler {
	return &AccountHandler{
		Usecase:   uc,
		Presenter: pr,
	}
}

func (u *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &AccountJsonView{w: w, successCode: http.StatusOK}
	accounts, err := u.Usecase.List(ctx)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.ViewModels(u.Presenter.ToViewModels(accounts))
}

func (u *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountUUID := chi.URLParam(r, "accountUUID")
	view := &AccountJsonView{w: w, successCode: http.StatusOK}
	account, err := u.Usecase.Find(ctx, accountUUID)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.View(u.Presenter.ToViewModel(account))
}

func (u *AccountHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &AccountJsonView{w: w, successCode: http.StatusCreated}
	in := &usecase.AccountStoreInput{}
	if err := json.NewDecoder(r.Body).Decode(in); err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	account, err := u.Usecase.Store(ctx, in)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.View(u.Presenter.ToViewModel(account))
}

func (u *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &AccountJsonView{w: w, successCode: http.StatusNoContent}
	_, err := u.Usecase.Delete(ctx, chi.URLParam(r, "accountUUID"))
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.ViewNoBody()
}
