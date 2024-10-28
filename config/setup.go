package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB configura e retorna a conexão com o banco de dados
func ConnectDB() *sql.DB {
	// Carrega as variáveis de ambiente, se ainda não foram carregadas
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado ou já carregado")
	}

	// Pega as variáveis do .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Configura a string de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Testa a conexão para garantir que está funcionando
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso")
	return db
}
