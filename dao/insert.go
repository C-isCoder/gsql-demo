package dao

import (
	"fmt"
	"gsql-demo/database"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		stmt, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
		if err != nil {
			fmt.Printf("Insert Prepare error:%v\n", err.Error())
		}
		if stmt == nil {
			return
		}
		_, _ = stmt.Exec(name, city)
		http.Redirect(w, r, "/", 301)
	}
}
