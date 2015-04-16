package database

import (
	"database/sql"
	"errors"
)

type DataBase struct {
	DataBaser
	driverName     string
	dataSourceName string
	DB             *sql.DB
}

func NewDataBase(driverName, dataSourceName string) *DataBase {
	return &DataBase{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}
}
func Open(driverName string, dataSourceName string) (*sql.DB, error) {
	if len(driverName) == 0 || len(dataSourceName) == 0 {
		return nil, errors.New("params is invalid")
	}

	return sql.Open(driverName, dataSourceName)
}

func (this *DataBase) Open() (*sql.DB, error) {
	if len(this.driverName) == 0 || len(this.dataSourceName) == 0 {
		return nil, errors.New("params is invalid")
	}

	DB, err := sql.Open(this.driverName, this.dataSourceName)
	this.DB = DB
	return DB, err
}

func (this *DataBase) Create(dbname string) error {
	_, err := this.DB.Exec("CREATE DATABASE " + dbname)
	return err
}

func (this *DataBase) Drop(dbname string) error {
	_, err := this.DB.Exec("DROP DATABASE " + dbname)
	return err
}

func (this *DataBase) Using(dbname string) error {
	this.Create(dbname)

	_, err := this.DB.Exec("use " + dbname)
	if err != nil {
		return err
	}

	return nil
}
