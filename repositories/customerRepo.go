package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/FawwazN/mnc-backend/api/config"
	"github.com/FawwazN/mnc-backend/api/models"
)

func SaveCustomers(customer models.Customer) error {
	customers, err := GetAllCustomers()
	if err != nil {
		return err
	}

	customers = append(customers, customer)

	data, err := json.MarshalIndent(customers, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.CustomerDB, data, 0644)
}

func GetAllCustomers() ([]models.Customer, error) {
	data, err := os.ReadFile(config.CustomerDB)
	if err != nil {
		return nil, err
	}

	var customers []models.Customer
	err = json.Unmarshal(data, &customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomerByUsername(username string) (models.Customer, error) {
	customers, err := GetAllCustomers()
	if err != nil {
		return models.Customer{}, err
	}

	for _, customer := range customers {
		if customer.Username == username {
			return customer, nil
		}
	}
	return models.Customer{}, errors.New("customer not found")
}

func GetCustomerById(username string) (models.Customer, error) {
	customers, err := GetAllCustomers()
	if err != nil {
		return models.Customer{}, err
	}

	for _, customer := range customers {
		if customer.Id == username {
			return customer, nil
		}
	}
	return models.Customer{}, errors.New("customer not found")
}

func UpdateCustomer(updatedCustomer models.Customer) error {
	customers, err := GetAllCustomers()
	if err != nil {
		return err
	}

	for i := range customers {
		if customers[i].Id == updatedCustomer.Id {
			customers[i] = updatedCustomer
			break
		}
	}

	data, err := json.MarshalIndent(customers, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(config.CustomerDB, data, 0644)
}
