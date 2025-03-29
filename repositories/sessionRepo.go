package repositories

import (
	"encoding/json"
	"os"

	"github.com/FawwazN/mnc-backend/api/config"
	"github.com/FawwazN/mnc-backend/api/models"
	"github.com/FawwazN/mnc-backend/api/utils"
)

func CreateSession(session models.Session) error {
	sessions, err := GetSessions()
	if err != nil {
		return err
	}

	sessions = append(sessions, session)

	data, err := json.MarshalIndent(sessions, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.SessionDB, data, 0644)
}

func GetSessions() ([]models.Session, error) {
	data, err := os.ReadFile(config.SessionDB)
	if err != nil {
		return nil, err
	}

	var sessions []models.Session
	err = json.Unmarshal(data, &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func FetchSession(token string) (models.Session, error) {
	sessions, err := GetSessions()
	if err != nil {
		return models.Session{}, err
	}

	for _, session := range sessions {
		if session.Token == token {
			return session, nil
		}
	}

	return models.Session{}, nil
}

func DeleteSession(token string) error {
	token = utils.ExtractToken(token)
	sessions, err := GetSessions()
	if err != nil {
		return err
	}

	var newSessions []models.Session
	for _, session := range sessions {
		if session.Token != token {
			newSessions = append(newSessions, session)
		}
	}

	data, err := json.MarshalIndent(newSessions, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(config.SessionDB, data, 0644)
}
