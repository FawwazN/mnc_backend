package services

import (
	"errors"
	"time"

	"github.com/FawwazN/mnc-backend/api/models"
	"github.com/FawwazN/mnc-backend/api/repositories"
	"github.com/google/uuid"
)

func PaymentService(username, merchantID string, amount int) (models.Transaction, error) {
	customer, err := repositories.GetCustomerByUsername(username)

	if err != nil {
		return models.Transaction{}, errors.New("customer not found")
	}

	merchant, err := repositories.GetMerchantById(merchantID)
	if err != nil {
		return models.Transaction{}, errors.New("merchant not found")
	}

	if customer.Balance < amount {
		return models.Transaction{}, errors.New("balance is not enough")
	}

	customer.Balance -= amount
	err = repositories.UpdateCustomer(customer)
	if err != nil {
		return models.Transaction{}, errors.New("failed to update customer balance. " + err.Error())
	}

	transaction := models.Transaction{
		Id:        uuid.New().String(),
		Merchant:  merchant,
		Customer:  customer,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	err = repositories.SaveTransaction(transaction)
	if err != nil {
		return models.Transaction{}, errors.New("failed to save transaction history. " + err.Error())
	}

	return transaction, nil
}
