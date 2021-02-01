package database

import (
	"curso-go/api-echo-orm/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectAndMigrate() (*gorm.DB, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() (*gorm.DB, error) {
	// TODO: mover credenciais e acesso para
	// a) ficheiro yaml.
	// b) variável de environment com o dsn completo.
	// c) usar thrid-party tipo o Consul.
	dsn := "user1:123456@tcp(127.0.0.1:33060)/booking"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// Migrate the schema
	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	return nil
}
