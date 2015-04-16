package model

import (
	"errors"
	"github.com/Congenital/orm/utils"
	"sync"
)

var Models map[string]*utils.Tcolumns = make(map[string]*utils.Tcolumns)
var tableLock sync.RWMutex

func Register(table interface{}) error {
	name, err := utils.GetTableName(table)
	if err != nil {
		return err
	}

	tableLock.RLock()
	defer tableLock.RUnlock()

	if _, ok := Models[name]; ok {
		return errors.New("table exists")
	}

	column, err := utils.GetColumns(table)
	if err != nil {
		return err
	}

	Models[name] = column
	return nil
}

func Unregister(table interface{}) error {
	name, err := utils.GetTableName(table)
	if err != nil {
		return err
	}

	tableLock.Lock()
	defer tableLock.Unlock()

	if _, ok := Models[name]; ok {
		delete(Models, name)
		return nil
	}

	return errors.New("Unregister table error")
}

func Get(table interface{}) *utils.Tcolumns {
	name, err := utils.GetTableName(table)
	if err != nil {
		return nil
	}

	tableLock.RLock()
	defer tableLock.RUnlock()

	if v, ok := Models[name]; ok {
		return v
	}

	return nil
}
