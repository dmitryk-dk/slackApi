package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dmitryk-dk/slackbot/models"
	"io/ioutil"
	"net/http"
	"errors"
	"net/url"
)

type ApiResponse struct {
	OK       bool   `json:"ok"`
	URL      string `json:"url"`
	Error    string `json:"error"`
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`

	Team  map[string]string `json:"team"`
	Self  map[string]string `json:"self"`

	Channel struct {
		Id string `json:"id"`
	} `json:"channel"`

	Channels []models.Channel `json:"channel"`
	Ims      []models.IM      `json:"ims"`
}

func (api ApiResponse) ReadConfig() (*string, *string) {
	flagName := api.ReadFlags()
	file, err := ioutil.ReadFile(*flagName)
	err = json.Unmarshal(file, &api)
	if err != nil {
		return nil, nil
	}
	return &api.Token, &api.Endpoint
}

func (api ApiResponse) ReadFlags() (configFile *string) {
	configFile = flag.String("config", "config.json", "Path to config file")
	flag.Parse()
	return
}

func SlackConnect() (*ApiResponse, error) {
	var api ApiResponse
	token, endpoint := api.ReadConfig()
	api = ApiResponse{
		Endpoint: *endpoint,
		Token: *token,
	}
	getUrl := api.Endpoint + api.Token
	resp, err := http.Get(getUrl)

	if err != nil {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &api)

	return &api, nil
}

func Action (action string, params map[string]string) (*ApiResponse,error){
	resp := &ApiResponse{}
	body, err := GetWrapper(action, params)
	if err != nil {
		return resp, err
	}

	json.Unmarshal(body, &resp)

	if !resp.OK {
		return resp, errors.New("slack error: " + resp.Error)
	}

	return resp, nil
}

func GetWrapper (action string, params map[string]string) ([]byte, error) {
	apiResp := &ApiResponse{}
	route, err := url.Parse(apiResp.Endpoint + action)

	if err != nil {
		return []byte{}, err
	}

	query := url.Values{}
	query.Add("token", apiResp.Token)

	for key, value := range params {
		query.Add(key, value)
	}

	route.RawQuery = query.Encode()

	response, err := http.Get(route.String())
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
