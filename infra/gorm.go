package external

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mini-rest-api/common"
	"mini-rest-api/common/command"
	"mini-rest-api/common/exception"
	"mini-rest-api/domain"
)

var dbInstance *gorm.DB

func NewGormDB() *gorm.DB {
	dsn := command.PostgreConfig()

	if dbInstance == nil {
		var err error

		DBInstance, err := gorm.Open(postgres.New(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		if common.IsMigrate && dbInstance == nil {
			err = DBInstance.AutoMigrate(&domain.User{}, &domain.Note{})
			exception.PanicIfNeeded(err)

			if err == nil {
				fmt.Println("Migrasi berhasil!")
			}
		}

		dbInstance = DBInstance

		return dbInstance
	} else {
		return dbInstance
	}
}
