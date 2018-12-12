package web

import (
	"net/http"

	"github.com/miya-masa/go-tx-sandbox/domain/entity"
)

type DepartmentJsonView struct {
	w           http.ResponseWriter
	successCode int
}

func (j *DepartmentJsonView) ErrorView(err error) {
	logger.Println(err)
	j.w.WriteHeader(http.StatusInternalServerError)
}

func (j *DepartmentJsonView) View(department *entity.Department) {
	if err := jsonView(j.w, j.successCode, department); err != nil {
		logger.Println(err)
		j.w.WriteHeader(http.StatusInternalServerError)
	}
}
