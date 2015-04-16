package table

import (
	"database/sql"
	"errors"
	"github.com/Congenital/orm/model"
	"github.com/Congenital/orm/utils"
)

type Table struct {
	DB *sql.DB
}

func NewTable(DB *sql.DB) *Table {
	return &Table{
		DB: DB,
	}
}

func (this *Table) Register(table interface{}) error {
	return model.Register(table)
}

func (this *Table) Unregister(table interface{}) error {
	return model.Unregister(table)
}

func (this *Table) Get(table interface{}) *utils.Tcolumns {
	return model.Get(table)
}

func (this *Table) Create(table interface{}) error {
	return create(this.DB, table, model.Get(table))
}

func (this *Table) Drop(table interface{}) error {
	return Drop(this.DB, table)
}

func (this *Table) Clear(table interface{}) error {
	return Clear(this.DB, table)
}

func (this *Table) Query(table interface{}) error {
	return nil
}

func (this *Table) QueryRow(table interface{}) error {
	return nil
}

func (this *Table) Exec(table interface{}) error {
	return nil
}

func Register(table interface{}) error {
	return model.Register(table)
}

func Unregister(table interface{}) error {
	return model.Unregister(table)
}

func Get(table interface{}) *utils.Tcolumns {
	return model.Get(table)
}

func create(DB *sql.DB, table interface{}, tables *utils.Tcolumns) error {
	var err error

	if DB == nil {
		return errors.New("Err - DB is nil")
	}

	column := tables
	if column == nil {
		column, err = utils.GetColumns(table)
		if err != nil {
			return err
		}
	}

	var sql string
	sql = "create table " + column.Name + "("
	for k, v := range column.Columns {
		sql += k + " " + utils.SqlTypeFormat(v)
		tag := column.Tags[k]

		for tagKey, tagValue := range tag {
			if tagValue == "true" {
				constraint := utils.GetConstraint(tagKey)
				sql = sql + " " + constraint
			}
		}
		sql = sql + ","
	}

	sql = sql[:len(sql)-1] + ")"

	_, err = DB.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Create(DB *sql.DB, table interface{}) error {
	return create(DB, table, nil)
}

func Drop(DB *sql.DB, table interface{}) error {
	if DB == nil {
		return errors.New("Err - DB is nil")
	}

	name, err := utils.GetTableName(table)
	if err != nil {
		return err
	}

	_, err = DB.Exec("DROP TABLE " + name)
	if err != nil {
		return err
	}

	return nil
}

func Clear(DB *sql.DB, table interface{}) error {
	if DB == nil {
		return errors.New("Err - DB is nil")
	}

	name, err := utils.GetTableName(table)
	if err != nil {
		return err
	}

	_, err = DB.Exec("DELETE FROM TABLE " + name)
	if err != nil {
		return err
	}

	return nil
}
