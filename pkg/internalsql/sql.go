package internalsql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName))

	if err != nil {
		log.Fatalf("error connect to database %+v\n", err)
	}

	return db, nil
}
