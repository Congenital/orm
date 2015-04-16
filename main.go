package main

import (
	"github.com/Congenital/log"
	"github.com/Congenital/orm/database"
	"github.com/Congenital/orm/table"
	_ "github.com/go-sql-driver/mysql"
)

type TestStruct struct {
	Name string `p:"true" f:"TestStruct2(Id)"`
	Age  int    `n:"true"`
}

type TestStruct2 struct {
	Id      int
	Address string
}

func main() {
	mydb := database.NewDataBase("mysql", "root:root@tcp(10.0.0.7:3306)/test")
	_, err := mydb.Open()
	if err != nil {
		log.Error(err)
		return
	}
	defer mydb.DB.Close()

	err = mydb.Create("Aaa")
	if err != nil {
		log.Error(err)
		return
	}
	defer mydb.Drop("Aaa")
	mydb.Using("Aaa")

	rela := table.NewTable(mydb.DB)
	err = rela.Register(TestStruct{})
	if err != nil {
		log.Error(err)
		return
	}
	defer rela.Unregister(TestStruct{})

	err = rela.Create(TestStruct{Name: "fdafdafasd", Age: 231})
	if err != nil {
		log.Error(err)
		return
	}
	defer rela.Drop(TestStruct{})
}
