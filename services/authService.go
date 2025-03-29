package services

import (
	"errors"

	"github.com/FawwazN/mnc-backend/api/models"
	"github.com/FawwazN/mnc-backend/api/repositories"
	"github.com/FawwazN/mnc-backend/api/utils"
)

func Login(username, password string) (map[string]interface{}, error) {
	customer, err := repositories.GetCustomerByUsername(username)

	if err != nil || customer.Password != password {
		return map[string]interface{}{}, errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(customer.Username)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = repositories.CreateSession(models.Session{
		Username:  customer.Username,
		Token:     token["token"].(string),
		ExpiredAt: token["expired_at"].(string),
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	return token, nil
}

func Logout(token string) (string, error) {
	onlyToken := utils.ExtractToken(token)

	session, err := repositories.FetchSession(onlyToken)
	if err != nil || session.Token == "" {
		return "", errors.New("invalid token")
	}

	validateResult, err := utils.VerifyToken(token)
	if err != nil {
		return "", err
	}

	err = repositories.DeleteSession(validateResult.Token)
	if err != nil {
		return "", err
	}

	return validateResult.Token, nil
}
