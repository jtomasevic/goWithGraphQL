package evax

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

)

type DataSource interface {
	DB() *gorm.DB
}

type DataSourceImpl struct {
	db *gorm.DB
}

func (dataSource *DataSourceImpl) DB() *gorm.DB {
	if dataSource.db == nil {
		db, err := gorm.Open(sqlite.Open("evax.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		fmt.Println("initialize evax gorm connection")
		dataSource.db = db
	}
	return dataSource.db
}


func NewEvaxDataSource() *DataSourceImpl {
	return &DataSourceImpl{
	}
}