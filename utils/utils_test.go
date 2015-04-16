package utils

import (
	"github.com/Congenital/log"
	"testing"
)

type T struct {
	Name string
}

func TestUtils(t *testing.T) {
	s := struct {
		Name string `p:"true"`
		Id   int    `n:"true"`
	}{}

	column, err := GetColumns(s)
	if err != nil {
		log.Error(err)
	}

	name, err := GetTableName(T{})
	if err != nil {
		log.Error(err)
		return
	}
}
