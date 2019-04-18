package dao

import (
	"fmt"
	"gsql-demo/database"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		stmt, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
		if err != nil {
			fmt.Printf("Update Prepare error:%v\n", err.Error())
		}
		if stmt == nil {
			return
		}
		_, _ = stmt.Exec(name, city, id)
		http.Redirect(w, r, "/", 301)
	}
}