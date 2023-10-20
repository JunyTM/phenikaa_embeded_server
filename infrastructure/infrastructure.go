package infrastructure

import (
	"flag"
	"log"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")
	flag.Parse()

	log.Println("Initializing database...")
	if err := InitDatabase(initDB); err != nil {
		log.Println("error initialize database: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func GetDBName() string {
	return "phenikaa_embedded"
}

func GetAppPort() string {
	return "12002"
}
