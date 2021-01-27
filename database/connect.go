package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// TODO: mover credenciais e accesso pata
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
