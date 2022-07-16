package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
	err error
)

type Driver string

const (
	MySQL Driver = "MYSQL"
)

//this is for migrate to databse mysql
func NewMySQL(){
	once.Do(func() {
		dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
		
		//open database
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MySQL database")
	})
}

func DB() *gorm.DB {
	return db
}
