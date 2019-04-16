package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "congxiaodan"
	dbName := "db_gsql"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Printf("open db error:%v\n", err.Error())
	}
	return db
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error:%v\n", err.Error())
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		fmt.Printf("index query error:%v\n", err.Error())
	}
	if rows == nil {
		return
	}
	emp := Employee{}
	var res []Employee
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
	_ = tmpl.ExecuteTemplate(w, "Index", res)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	nId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		fmt.Printf("Show query error:%v\n", err.Error())
	}
	if rows == nil {
		return
	}
	emp := Employee{}
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
	_ = tmpl.ExecuteTemplate(w, "Show", emp)
}

func New(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	nId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
	if err != nil {
		fmt.Printf("edit query error:%v\n", err.Error())
	}
	if rows == nil {
		return
	}
	emp := Employee{}
	for rows.Next() {
		var id int
		var name, city string
		err = rows.Scan(&id, &name, &city)
		if err != nil {
			fmt.Printf("edit scan error:%v\n", err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	_ = tmpl.ExecuteTemplate(w, "Edit", emp)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		stmt, err := db.Prepare("INSERT INTO employee(name,city) VALUES(?,?)")
		if err != nil {
			fmt.Printf("Insert Prepare error:%v\n", err.Error())
		}
		if stmt == nil {
			return
		}
		_, _ = stmt.Exec(name, city, id)
		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
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

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
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
