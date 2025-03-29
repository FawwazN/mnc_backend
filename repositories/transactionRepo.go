package repositories

import (
	"encoding/json"
	"os"

	"github.com/FawwazN/mnc-backend/api/config"
	"github.com/FawwazN/mnc-backend/api/models"
)

func SaveTransaction(history models.Transaction) error {
	transactions, err := GetAllTransaction()
	if err != nil {
		return err
	}

	transactions = append(transactions, history)

	data, err := json.MarshalIndent(transactions, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.TransactionDB, data, 0644)
}

func GetAllTransaction() ([]models.Transaction, error) {
	data, err := os.ReadFile(config.TransactionDB)
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
