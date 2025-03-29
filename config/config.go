package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerPort    string
	JWTSecret     string
	CustomerDB    string
	MerchantDB    string
	TransactionDB string
	SessionDB     string
)

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	ServerPort = getEnv("SERVER_PORT", "8080")
	JWTSecret = getEnv("JWT_SECRET", "thisissecretcode")
	CustomerDB = getEnv("CUSTOMER_DB", "storage/customers.json")
	MerchantDB = getEnv("MERCHANT_DB", "storage/merchants.json")
	TransactionDB = getEnv("TRANSACTION_DB", "storage/transactions.json")
	SessionDB = getEnv("SESSION_DB", "storage/sessions.json")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
