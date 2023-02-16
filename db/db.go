package db

import (
	"fmt"
	"log"
	"os"

	"github.com/guilherm5/crud/handlers"
	"github.com/guilherm5/crud/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao localizar arquivo de configuração")
	}

	Host := os.Getenv("HOST")
	DBPort := os.Getenv("DBPORT")
	User := os.Getenv("USER")
	Password := os.Getenv("PASSWORD")
	DBName := os.Getenv("DBNAME")

	stringConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, DBPort, User, Password, DBName)

	db, err := gorm.Open(postgres.Open(stringConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao realizar conexão")
	} else {
		fmt.Println("Sucesso ao realizar conexão")
	}

	db.AutoMigrate(&models.Tarefas{})

	return db
}

func LoadDB(DB *gorm.DB) handlers.Handler {
	return handlers.Handler{DB: DB}
}
