package api

import "github.com/dmitryk-dk/slackbot/models"

func ListChannels(token string) ([]models.Channel, error) {
	resp, err := Action("channels.list", map[string]string{})
	if err != nil {
		return []models.Channel{}, err
	}

	return resp.Channels, nil
}
