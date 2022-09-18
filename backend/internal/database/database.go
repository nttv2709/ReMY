package database

import (
	"context"
	"fmt"
	"log"
	"remy/pkg/ent"
)

func Init() (*ent.Client, error) {
	// Connect database
	DB_HOST := "ec2-34-200-205-45.compute-1.amazonaws.com"
	DB_OUT_PORT := "5432"
	DB_IN_PORT := "5432"
	DB_USER := "ethidrlsvjpwss"
	DB_PASS := "7d04e57ffe09ea310e3838082ae4a4fa17948af82fb690e73a203780eccb5ab4"
	DB_NAME := "df6687mil3q4ar"
	fmt.Printf("db_out: %s, db_in: %s\n", DB_OUT_PORT, DB_IN_PORT)
	fmt.Printf("Connecting to Mysql: %s:%s@tcp(%s:%s)/%s?parseTime=True\n", DB_USER, DB_PASS, DB_HOST, DB_OUT_PORT, DB_NAME)
	// postgres://ethidrlsvjpwss:7d04e57ffe09ea310e3838082ae4a4fa17948af82fb690e73a203780eccb5ab4@ec2-34-200-205-45.compute-1.amazonaws.com:5432/df6687mil3q4ar
	database, errDB := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DB_HOST, DB_OUT_PORT, DB_USER, DB_NAME, DB_PASS))
	fmt.Println("Connect to Database successfully!")
	if errDB != nil {
		return nil, errDB
	}
	// Run the auto migration tool.
	if errSchema := database.Schema.Create(context.Background()); errSchema != nil {
		log.Fatalf("failed creating schema resources: %v", errSchema)
		log.Fatalf("failed creating schema resources: %v", errSchema)
		return nil, errSchema
	}
	return database, nil
}
