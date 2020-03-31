package query

import (
	"github.com/Snehal1112/QueryBuilder/ddl"
	"github.com/Snehal1112/QueryBuilder/dml"
)

type SQL interface {
	NewDDL() ddl.Service
	NewDML() dml.Service
	dataSourceName(driver string) string
	SelectDB(name string) error
	GetSelectedDB() string
}