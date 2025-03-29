package services

import (
	"errors"

	"github.com/FawwazN/mnc-backend/api/models"
	"github.com/FawwazN/mnc-backend/api/repositories"
)

func CreateCustomerService(customer models.Customer) error {
	existingCustomer, err := repositories.GetCustomerByUsername(customer.Username)
	if err == nil && existingCustomer.Id != "" {
		return errors.New("username already exists")
	}

	return repositories.SaveCustomers(customer)
}
func GetCustomerByIdService(id string) (models.Customer, error) {
	return repositories.GetCustomerById(id)
}
func GetAllCustomersService() ([]models.Customer, error) {
	return repositories.GetAllCustomers()
}
func GetCustomerByUsernameService(username string) (models.Customer, error) {
	return repositories.GetCustomerByUsername(username)
}
func UpdateCustomerService(customer models.Customer) error {
	_, err := repositories.GetCustomerById(customer.Id)
	if err != nil {
		return errors.New("customer not found")
	}
	return repositories.UpdateCustomer(customer)
}
