package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/FawwazN/mnc-backend/api/config"
	"github.com/FawwazN/mnc-backend/api/models"
)

func SaveMerchants(merchant models.Merchant) error {
	merchants, err := GetAllMerchants()
	if err != nil {
		return err
	}

	merchants = append(merchants, merchant)

	data, err := json.MarshalIndent(merchants, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.MerchantDB, data, 0644)
}

func GetAllMerchants() ([]models.Merchant, error) {
	data, err := os.ReadFile(config.MerchantDB)
	if err != nil {
		return nil, err
	}

	var merchants []models.Merchant
	err = json.Unmarshal(data, &merchants)
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func GetMerchantById(username string) (models.Merchant, error) {
	merchants, err := GetAllMerchants()
	if err != nil {
		return models.Merchant{}, err
	}

	for _, merchant := range merchants {
		if merchant.Id == username {
			return merchant, nil
		}
	}
	return models.Merchant{}, errors.New("merchant not found")
}

func UpdateMerchant(merchant models.Merchant) error {
	merchants, err := GetAllMerchants()
	if err != nil {
		return err
	}

	for i, mrc := range merchants {
		if mrc.Id == merchant.Id {
			merchants[i] = merchant
			break
		}
	}
	data, err := json.MarshalIndent(merchants, "", "		")
	if err != nil {
		return err
	}

	return os.WriteFile(config.MerchantDB, data, 0644)
}
