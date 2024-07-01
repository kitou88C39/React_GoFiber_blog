package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DBConn *gorm.DB

func main() {
	
	dsn := "root:neeraj@tcp(127.0.0.1:3306)/fiber_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Defa.Default.LogMode(logger.Error),

	})
  }