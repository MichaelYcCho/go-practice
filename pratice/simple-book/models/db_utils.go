package models

import (
	"gorm.io/gorm"
)

func GetTableList() []string {
	println("GetTableList Start")

	TableList, _ := DB.Debug().Migrator().GetTables()

	return TableList
}

// 사용하지 model struct에서 사용하지 않는 column을 삭제하는
func DropUnusedColumns(table any) {

	stmt := &gorm.Statement{DB: DB}

	stmt.Parse(table)

	fields := stmt.Schema.Fields
	columns, _ := DB.Debug().Migrator().ColumnTypes(table)

	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			DB.Migrator().DropColumn(table, columns[i].Name())
		}
	}
}
