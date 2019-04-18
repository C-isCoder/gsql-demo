package dao

import (
	"gsql-demo/form"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	_ = form.Temp.ExecuteTemplate(w, "New", nil)
}
