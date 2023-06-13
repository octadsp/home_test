package database

import (
	"fmt"

	"github.com/octadsp/server-home-test/models"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
)

func RunMigration() {
	err := connectiondb.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}
	fmt.Println("Migration Success!")
}
