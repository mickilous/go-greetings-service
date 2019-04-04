package deserve

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
	logger     *log.Logger
	httpClient *http.Client
	srvAddr    string
}

type Message struct {
	Deserve string `json:"deserve"`
}

func NewClient(logger *log.Logger, client *http.Client, srvAddr string) *Client {
	return &Client{
		logger:     logger,
		httpClient: client,
		srvAddr:    srvAddr,
	}
}

func (c *Client) IsGreetable(userId string) bool {

	const endPoint = "/deserve/"
	const defaultRet = true //In doubt let's greet and do not be impolite

	response, err := c.httpClient.Get(c.srvAddr + endPoint + userId)
	defer response.Body.Close()
	if err != nil {
		c.logger.Printf("The HTTP request failed with error %v\n", err)
		return defaultRet
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.logger.Printf("The HTTP body read failed with error %v\n", err)
		return defaultRet
	}
	if response.StatusCode != http.StatusOK {
		c.logger.Printf("The HTTP request failed with HTTP Code : %v - %v", response.StatusCode, string(body))
		return defaultRet
	}

	deserveMsg := &Message{}
	json.Unmarshal(body, &deserveMsg)
	greetable, _ := strconv.ParseBool(deserveMsg.Deserve)
	return greetable

}
