package client

import (
	"net/http"

	"github.com/dmitryk-dk/slackbot/api"
	"github.com/dmitryk-dk/slackbot/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Token     string
	Connected bool
	Incoming  chan interface{}
	Errors    chan error
	conn      *websocket.Conn

	messageID int
}

func CreateClient(token string) *Client {
	return &Client{
		Token:     token,
		Incoming:  make(chan interface{}),
		Errors:    make(chan error, 1),
		messageID: 1,
	}
}

func NewClient(token string) *Client {
	return CreateClient(token)
}

func (c *Client) Connect(url string) error {

	dialer := websocket.Dialer{}
	headers := http.Header{}

	connection, _, err := dialer.Dial(url, headers)
	if err != nil {
		return err
	}

	c.conn = connection
	c.Connected = true
	return nil
}

func (c *Client) Listen() {
	connection := c.conn

	go func() {
		for {
			_, msg, err := connection.ReadMessage()
			if err != nil {
				c.Errors <- err
				return
			}

			message, err := api.EventParser(msg)
			if err != nil {
				c.Errors <- err
				return
			}

			c.Incoming <- message
		}
	}()
}

func (c *Client) Loop() {
	c.Listen()
}

func (c *Client) OpenIm(user string) (string, error) {
	return api.OpenIm(c.Token, user)
}

func (c *Client) SendMessage(channel, text string) error {
	msg := struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"text"`
	}{c.messageID, "message", channel, text}

	if err := c.conn.WriteJSON(msg); err != nil {
		return err
	}

	c.messageID++

	return nil
}

func (c *Client) ListChannels() ([]models.Channel, error) {
	return api.ListChannels(c.Token)
}

func (c *Client) ListIms() ([]models.IM, error) {
	return api.ListIms(c.Token)
}

func (c *Client) Ping() error {
	msg := struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	}{c.messageID, "ping"}

	if err := c.conn.WriteJSON(msg); err != nil {
		return err
	}

	c.messageID++

	return nil
}
