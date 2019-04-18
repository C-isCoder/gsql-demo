package dao

import (
	"fmt"
	"gsql-demo/database"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()
	id := r.URL.Query().Get("id")
	stmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		fmt.Printf("Delete Prepare error:%v\n", err.Error())
	}
	if stmt == nil {
		return
	}
	_, _ = stmt.Exec(id)
	http.Redirect(w, r, "/", 301)
}
