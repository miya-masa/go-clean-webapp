package web

import (
	"net/http"

	"github.com/miya-masa/go-tx-sandbox/domain/usecase"
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
