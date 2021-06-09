package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Clients struct {
	gorm.Model
	Age  string
	Name string
	City string
}

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/your-db-name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Clients{})

	csvFile, err := os.Open("lista.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		db.Create(&Clients{Name: line[0], Age: line[1], City: line[2]})
	}
}
