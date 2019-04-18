package dao

import (
	"fmt"
	"gsql-demo/database"
	"gsql-demo/form"
	"gsql-demo/model"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()
	nId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		fmt.Printf("Show query error:%v\n", err.Error())
	}
	if rows == nil {
		return
	}
	emp := model.Employee{}
	for rows.Next() {
		var id int
		var name, city string
		err = rows.Scan(&id, &name, &city)
		if err != nil {
			fmt.Printf("Show scan error:%v\n", err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	_ = form.Temp.ExecuteTemplate(w, "Show", emp)
}
