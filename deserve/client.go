package deserve

import (
	"encoding/json"
	"greetings-service/configuration"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
	*log.Logger
	*http.Client
	ServerAddr string
}

func NewClient(provider configuration.Provider, logger *log.Logger, client *http.Client) *Client {
	return &Client{
		logger,
		client,
		provider.GetString("DESERVE_ADDR", "http://localhost:8090")}
}

type Message struct {
	Deserve string `json:"deserve"`
}

func (c *Client) IsGreetable(userId string) bool {

	const endPoint = "/deserve/"
	const defaultRet = true //In doubt let's greet and do not be impolite

	response, err := c.Client.Get(c.ServerAddr + endPoint + userId)
	if err != nil {
		c.Logger.Printf("The HTTP request failed with error %v\n", err)
		return defaultRet
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.Logger.Printf("The HTTP body read failed with error %v\n", err)
		return defaultRet
	}
	if response.StatusCode != http.StatusOK {
		c.Logger.Printf("The HTTP request failed with HTTP Code : %v - %v", response.StatusCode, string(body))
		return defaultRet
	}

	deserveMsg := &Message{}
	json.Unmarshal(body, &deserveMsg)
	greetable, _ := strconv.ParseBool(deserveMsg.Deserve)
	return greetable

}
