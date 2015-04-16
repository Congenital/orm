package utils

import (
	"errors"
	"github.com/Congenital/orm/contrast"
	"reflect"
	"strings"
)

type Tcolumns struct {
	Name    string
	Columns map[string]string
	Tags    map[string]map[string]string
	Values  map[string]interface{}
}

func GetColumns(table interface{}) (*Tcolumns, error) {
	tp := reflect.ValueOf(table)
	tpType := tp.Type()
	if tp.Kind() != reflect.Struct {
		return nil, errors.New("tp not struct")
	}

	column := &Tcolumns{}
	column.Name = strings.ToLower(tpType.Name())
	column.Columns = make(map[string]string)
	column.Tags = make(map[string]map[string]string)
	column.Values = make(map[string]interface{})

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tpTmp := field.Type()
		fieldname := strings.ToLower(tpType.Field(i).Name)

		column.Values[fieldname] = field.Interface()
		column.Columns[fieldname] = strings.ToLower(tpTmp.Kind().String())

		if tpType.Field(i).Tag == "" {
			continue
		}

		column.Tags[fieldname] = make(map[string]string)

		for tagNum := 0; tagNum < contrast.Tags_length; tagNum++ {
			tag := strings.ToLower(tpType.Field(i).Tag.Get(contrast.Tags[tagNum]))
			if tag != "" {
				column.Tags[fieldname][strings.ToLower(contrast.Tags[tagNum])] = tag
			}
		}
	}

	return column, nil
}

func SqlTypeFormat(Type string) string {
	return contrast.SqlType[Type]
}

func GetConstraint(constraint string) string {
	return contrast.Constraint[constraint]
}

func GetTableName(table interface{}) (string, error) {
	tp := reflect.TypeOf(table)
	if tp.Kind() != reflect.Struct {
		return "", errors.New("no struct")
	}

	return tp.Name(), nil
}
