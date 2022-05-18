package migrations

import (
	"fmt"

	ds "github.com/evax/app/data_sources/evax"
	task "github.com/evax/app/services/tasks/repository"
	user "github.com/evax/app/services/users/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var clearUsers = true
var generateSampleData = false

func Strptr(value string) *string {
	return &value
}

func EvaxDbAutoMigrate() {
	db, err := gorm.Open(sqlite.Open("evax.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(user.User{})
	db.AutoMigrate(task.Task{})

	if generateSampleData {
		clearUsers = true
	}

	if clearUsers {
		db.Where("ID LIKE ?", "%").Delete(task.Task{})
		db.Where("ID LIKE ?", "%").Delete(user.User{})
	}
	dataSource := ds.NewEvaxDataSource()

	userRepo := user.NewUserRepository(dataSource)

	if generateSampleData {
		fmt.Println("******** generate users ***********")
		u1 := user.User{
			FirstName: Strptr("Igor"),
			LastName:  Strptr("Lazetic"),
			Username: "igorlaz",
			Password: "igorovasifra",
		}
		userRepo.SaveUser(u1)
		u2 := user.User{
			FirstName: Strptr("Marta"),
			LastName:  Strptr("Mirkovic"),
			Username: "martami",
			Password: "martinasifra",
		}
		userRepo.SaveUser(u2)
	}


}
