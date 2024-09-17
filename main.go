package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	CPF  uint32
	Id_address uint16
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
}

type Product struct {
	gorm.Model
	Code  string
	Price float32
	Name  string
	NCM   string
	UM    string
	Description string
	Id_Class uint8
}

type Class struct {
	gorm.Model
	Name string
	Description string
}

func main()  {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=America/Sao_Paulo",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}


