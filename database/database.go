package database

import (
	"fmt"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialiseDatabase(dbName string) {
	var (
		databaseUser     string = utils.GetValueFromConfigFile("DB_USER")
		databasePassword string = utils.GetValueFromConfigFile("DB_PASSWORD")
		databaseHost     string = utils.GetValueFromConfigFile("DB_HOST")
		databasePort     string = utils.GetValueFromConfigFile("DB_PORT")
		databaseName     string = dbName
	)

	// data source for MySQL
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	// variable to store the error
	var err error

	// create a connection to the database
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB Connection Established Successfully")

	err = DB.AutoMigrate(&models.User{}, &models.Item{})
	if err != nil {
		return
	}

	fmt.Println("DB Migration Complete")
}
