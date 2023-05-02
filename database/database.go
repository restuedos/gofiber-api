package database

import (
	"errors"
	"fmt"
	"strconv"

	"gofiber-example/config"
	"gofiber-example/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// connectDB
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	if err = db.AutoMigrate(&model.User{}); err == nil && db.Migrator().HasTable(&model.User{}) {
		if err := db.First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := hashPassword("Edo998877!")
			if err != nil {
				panic("Couldn't hash password for database seeding")
			}

			db.Create(&model.User{
				Email:    "restuedosetiaji@gmail.com",
				Username: "EDZero",
				Name:     "Restu Edo Setiaji",
				Password: hash,
			})

			fmt.Println("Database Seeded")
		}
	}
	db.AutoMigrate(&model.Book{})

	fmt.Println("Database Migrated")

	DBConn = db
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
