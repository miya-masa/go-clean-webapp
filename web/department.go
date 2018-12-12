package web

import (
	"net/http"

	"github.com/miya-masa/go-clean-webapp/usecase"
)

type DepartmentHandler struct {
	Usecase usecase.DepartmentInputPort
}

func (d *DepartmentHandler) Get(w http.ResponseWriter, r *http.Request) {
}

func (d *DepartmentHandler) Post(w http.ResponseWriter, r *http.Request) {
}

func (d *DepartmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
