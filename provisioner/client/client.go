package client

import (
	"../structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	host string
	port int
	username string
	password string
	resourceType string
	httpClient *http.Client
}

func NewClient(host string, port int, username string, password string, resourceType string) *Client {
	return &Client{
		host:host,
		port:port,
		username:username,
		password:password,
		resourceType:resourceType,
		httpClient:&http.Client{},
	}
}

// GetAll Retrieves all of the Items from the server
func (c *Client) GetAll()  {
	log.Println("Calling GetAll from client")
}

// GetItem gets an item with a specific name from the server
func (c *Client) GetItem(resourcePath string,name string) (*structs.Policy, error) {

	log.Println("Calling GetItem from client")
	body, err := c.httpRequest(fmt.Sprintf("%s/%v",resourcePath, name), "GET", bytes.Buffer{})
	if err != nil {
		log.Println("Inside GetItem error ",err)
		return nil, err
	}
	policy := &structs.Policy{}
	err = json.NewDecoder(body).Decode(policy)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// NewItem creates a new Item
func (c *Client) NewItem(resourcePath string, policy structs.Policy) error {

	log.Println("Calling NewItem from client")
	log.Println(c.requestPath(resourcePath))

	buff:=bytes.Buffer{}
	err:=json.NewEncoder(&buff).Encode(policy)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf(resourcePath), "POST", buff)
	return nil
}

// UpdateItem updates the values of an item
func (c *Client) UpdateItem() error {

	log.Println("Calling UpdateItem from client")

	return nil
}

// DeleteItem removes an item from the server
func (c *Client) DeleteItem() error {

	log.Println("Calling DeleteItem from client")
	return nil
}

func (c *Client) httpRequest(path,method string, body bytes.Buffer) (closer io.ReadCloser, err error)  {
	log.Println("httpRequest ",c.requestPath(path))


	req, err := http.NewRequest(method,c.requestPath(path),&body)
	if err != nil {

		log.Println("Error in creating req")
		log.Println(err)
		return nil, err
	}

	req.SetBasicAuth("admin","password")
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
		//req.Header.Add("Content-Type", "application/json")
	}

	log.Println("Calling Api...")
	log.Println(req)
	log.Println("Calling Api close...")


	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("Inside Error")
		log.Println(err)
		return nil, err
	}

	log.Println("response is:",resp)

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}


func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.host, c.port, path)
}
