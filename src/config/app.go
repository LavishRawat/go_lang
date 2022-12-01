package config

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

var (
	db *gorm.DB
)

const (
    HOST = "localhost"
    PORT = 5432
    USER = "postgres"     
	PASSWORD = "Lavish@123"
    DBNAME = "react_app"
 )

func Connect() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
   	)
	d, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
   	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB{
	return db
}