package dao

import (
	"fmt"
	"gsql-demo/database"
	"gsql-demo/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		fmt.Printf("index query error:%v\n", err.Error())
	}
	if rows == nil {
		return
	}
	emp := model.Employee{}
	var res []model.Employee
	for rows.Next() {
		var id int
		var name, city string
		err = rows.Scan(&id, &name, &city)
		if err != nil {
			fmt.Printf("scan query error:%v\n", err.Error())
		}
		emp.Id = id
		emp.City = city
		emp.Name = name
		res = append(res, emp)
	}
	_ = main.TMPL.ExecuteTemplate(w, "Index", res)
}
