package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gsql-demo/dao"
	"net/http"
)

func main() {
	http.HandleFunc("/", dao.Index)
	http.HandleFunc("/show", dao.Show)
	http.HandleFunc("/new", dao.New)
	http.HandleFunc("/edit", dao.Edit)
	http.HandleFunc("/insert", dao.Insert)
	http.HandleFunc("/update", dao.Update)
	http.HandleFunc("/delete", dao.Delete)
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error:%v\n", err.Error())
	}
}
