package api

import (
	"github.com/dmitryk-dk/slackbot/models"
)

func OpenIm (token, user string) (string, error) {
	resp, err := Action("im.open", map[string]string{"user": user})
	if err != nil {
		return "", err
	}

	return resp.Channel.Id, nil
}

func ListIms (token string) ([]models.IM, error) {
	resp, err := Action("im.list", map[string]string{})
	if err != nil {
		return []models.IM{}, err
	}

	return resp.Ims, nil
}
