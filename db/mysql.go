package db

import (
    "gin/basic/model"


    "gorm.io/driver/mysql"  // Import the MySQL driver package
    "gorm.io/gorm"
)
var MySQLDB *gorm.DB
func GetMySQLDBConnection() (*gorm.DB, error) {
    dsn := "user:1234@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    db.AutoMigrate(&model.Post{})
    db.AutoMigrate(&model.User{})
    MySQLDB = db  // Store the DB connection
    return db, err
}
