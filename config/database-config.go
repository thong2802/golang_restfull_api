package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetUpDatabaseConnection is creating a new connection to our database
func SetUpDatabaseConnection() *gorm.DB {
  errEvn:= godotenv.Load()
  if errEvn != nil {
	  panic("Fail to load evn file.")
  }

  dbUser := os.Getenv("DB_USER")
  dbPASS := os.Getenv("DB_PASS")
  dbHOST := os.Getenv("localhost")
  dbNAME := os.Getenv("golang_api")

  dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPASS,dbHOST, dbNAME)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
	  panic("Fail to create a connection to database")
  }

  //nanti kita isi modelnya di sini
  //db.AutoMigrate()
  return db

}
//CloseDatabaseConnection method is closing a connection between your app and database
func CloseDatabaseConnection(db *gorm.DB){
  dbSQL, err := db.DB()
  if err != nil {
	  panic("Fail to close connection to database")
  }
  dbSQL.Close()
}
