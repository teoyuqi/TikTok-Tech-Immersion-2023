package main

import (
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Message struct {
	rpc.Message
}

func initDB() *gorm.DB  {
	dsn := "user:password@tcp(host.docker.internal:3306)/db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(Message{})
	if err != nil {
		panic("failed to set up database")
	}

	return db
}
