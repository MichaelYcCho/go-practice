package database

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func GetTableList() []string {
	println("GetTableList Start")

	TableList, _ := DB.Debug().Migrator().GetTables()

	return TableList
}

// db autoMigrate with transaction
func AutoMigrateWithTransaction(db *gorm.DB, modelList []interface{}) error {

	defer func() {
		if r := recover(); r != nil {
			panic(errors.New("AutoMigrateWithTransaction Error"))
		}
	}()

	// raw SQL을 보고싶다면 Debug를 쓸것, db.Debug().AutoMigrate(modelList...);
	if err := db.AutoMigrate(modelList...); err != nil {
		log.Println("Migration Error: ", err)
		return err
	}

	return nil
}

// TODO: 최소 전체 struct 리스트를 받으면 해당 struct를 돌면서 자동으로 필드삭제라도 할수있게 흠....
// 사용하지 model struct에서 사용하지 않는 column을 삭제하는
func DropUnusedColumns(db *gorm.DB, modelList []interface{}) {

	// tx.Create(&user1)

	// tx.SavePoint("sp1")
	// tx.Create(&user2)
	// tx.RollbackTo("sp1") // Rollback user2

	// tx.Commit()

	stmt := &gorm.Statement{DB: db}

	for _, model := range modelList {
		println("시작")

		stmt.Parse(model)

		fields := stmt.Schema.Fields
		columns, _ := DB.Debug().Migrator().ColumnTypes(model)

		for i := range columns {
			found := false
			for j := range fields {
				if columns[i].Name() == fields[j].DBName {
					found = true
					break
				}
			}
			if !found {
				DB.Migrator().DropColumn(model, columns[i].Name())
			}
		}
	}
}

/*

// make modelList string with ,
	var modelListString string
	for i, model := range modelList {
		modelListString += fmt.Sprintf("%v", model)
		if i != length-1 {
			modelListString += ","
		}
	}
	log.Println("modelListString", modelListString)

*/
