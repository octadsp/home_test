package connectiondb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// dbCharset := os.Getenv("DB_CHARSET")
	// dbParseTime := os.Getenv("DB_PARSETIME")
	// dbLoc := os.Getenv("DB_LOC")

	// dsn := dbUser + ":@" + dbPass + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=" + dbParseTime + "&loc=" + dbLoc
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "root:@tcp(127.0.0.1:3306)/bosshire?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
