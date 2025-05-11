package database

import (
	"errors"
	"fmt"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/utils"
	"golang.org/x/crypto/bcrypt"
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

func CreateFakeItemsForTesting() (models.Item, error) {
	item, err := utils.CreateFaker[models.Item]()
	if err != nil {
		return models.Item{}, err
	}
	DB.Create(&item)
	fmt.Println("DB Test Item Created Successfully")
	return item, nil
}

func CreateFakeUsersForTesting() (models.User, error) {
	user, err := utils.CreateFaker[models.User]()
	if err != nil {
		return models.User{}, err
	}

	// create a password with bcrypt
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	// create the user
	var testUser models.User = models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: string(password),
	}
	DB.Create(&testUser)

	fmt.Println("DB Test User Created Successfully")

	return user, nil
}

func CleanTestData() {
	// remove all data inside the items table
	items := DB.Exec("TRUNCATE items")

	// remove all data inside the items table
	users := DB.Exec("TRUNCATE users")

	// check if the operation failed
	operationFailed := items.Error != nil || users.Error != nil

	if operationFailed {
		panic(errors.New("operation Failed. Cleaning Test Data Failed"))
	}

	fmt.Println("DB Test Items Cleaned Successfully")
}
