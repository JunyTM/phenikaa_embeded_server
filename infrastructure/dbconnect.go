package infrastructure

import (
	"embedded/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openConnection() (*gorm.DB, error) {
	connectSQL := "host=localhost user=postgres dbname=phenikaa_embedded port=5432 password=147563 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Printf("Not connect to database: %+v\n", err)
		return nil, err
	}
	return db, nil
}

func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		return err
	}

	if allowMigrate {
		log.Println("Migrating database...")
		db.Debug().AutoMigrate(
			&model.TrafficLight{},
		)
	}

	return nil
}
