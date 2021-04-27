package tenantapigowrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dogpakk/lib/mongolist"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	APIRootV1 = "scriptapi/v1"

	endpointList   = "get/%s"
	endpointSingle = "%s/%s"
	endpointCreate = "%s"
)

type Client struct {
	APIDomain string
	APISecret string
}

func (c Client) getEndpoint(rest string) string {
	return fmt.Sprintf("https://%s/%s/%s", c.APIDomain, APIRootV1, rest)
}

func (c Client) getListEndpoint(entityName string) string {
	return c.getEndpoint(fmt.Sprintf(endpointList, entityName))
}

func (c Client) getSingleEndpoint(entityName string, entityID primitive.ObjectID) string {
	return c.getEndpoint(fmt.Sprintf(endpointSingle, entityName, entityID.Hex()))
}

func (c Client) getCreateEndpoint(entityName string) string {
	return c.getEndpoint(fmt.Sprintf(endpointCreate, entityName))
}

func (c Client) GetEntityList(listEntity APIListEntity, listState mongolist.ListState) error {
	listStateBytes, err := json.Marshal(listState)
	if err != nil {
		return err
	}

	endpoint := c.getListEndpoint(listEntity.getEntityListName())

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(listStateBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APISecret))

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(listEntity); err != nil {
		return err
	}

	return nil
}

func (c Client) UpdateEntity(entity APISingleEntity, body []byte) error {
	endpoint := c.getSingleEndpoint(entity.getEntitySingleName(), entity.getID())

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APISecret))

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("received Non-OK response from Pakk")
	}

	if err := json.NewDecoder(resp.Body).Decode(entity); err != nil {
		return err
	}

	return nil
}
