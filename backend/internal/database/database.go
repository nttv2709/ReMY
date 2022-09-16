package database

import (
	"context"
	"fmt"
	"log"
	"sample_shecodes2022/pkg/ent"
)

var client *ent.Client

func Init() error {
	// Connect database
	DB_HOST := "localhost"
	DB_OUT_PORT := "13306"
	DB_IN_PORT := "3306"
	DB_USER := "root"
	DB_PASS := "P@ssw0rd"
	DB_NAME := "sample"
	fmt.Printf("db_out: %s, db_in: %s\n", DB_OUT_PORT, DB_IN_PORT)
	fmt.Printf("Connecting to Mysql: %s:%s@tcp(%s:%s)/%s?parseTime=True\n", DB_USER, DB_PASS, DB_HOST, DB_IN_PORT, DB_NAME)
	database, errDB := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_IN_PORT, DB_NAME))
	fmt.Println("Connect to Database successfully!")
	if errDB != nil {
		return errDB
	}
	// Run the auto migration tool.
	if errSchema := database.Schema.Create(context.Background()); errSchema != nil {
		log.Fatalf("failed creating schema resources: %v", errSchema)
	}
	client = database
	return nil
}
