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

	//httpClient := http.Client{
	//	Timeout: time.Second * 2, // Maximum of 2 secs
	//}

	const endPoint = "/deserve/"
	response, err := c.httpClient.Get(c.srvAddr + endPoint + userId)
	if err != nil {
		c.logger.Printf("The HTTP request failed with error %v\n", err)
		return true //In doubte let's greet and do not be impolite
	}
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		c.logger.Printf("The HTTP request failed with HTTP Code : %v - %v", response.StatusCode, string(body))
		return true //In doubte let's greet and do not be impolite
	}
	deserveMsg := &Message{}
	json.Unmarshal(body, &deserveMsg)
	greetable, _ := strconv.ParseBool(deserveMsg.Deserve)
	return greetable

}
