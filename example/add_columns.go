package main

import (
	"log"

	builder "github.com/Snehal1112/QueryBuilder"
	"github.com/Snehal1112/QueryBuilder/query"
)

func AddColumns(builder builder.SQL) {
	ddlQuery := builder.NewDDL()
	result, err := ddlQuery.Alter().Table("customers").Add().Column("age", query.NewDataType(query.INT,50), query.NewConstrain([]int{})).InsertAt(true,"name").Execute()
	if err != nil {
		log.Fatal(err)
	}
	_, err  = result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
}
