package tenantapigowrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	APIRootV1 = "pakkapi/v2"

	endpointList   = "get/%s"
	endpointSingle = "%s/%s"
	endpointCreate = "%s"
)

type Client struct {
	APIDomain string
	APIKey    string
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

func (c Client) GetEntityList(listEntity APIListEntity, listState ListSpec) error {
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
	req.Header.Set("Authorization", fmt.Sprintf("%s", c.APIKey))

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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read body to a string
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		bodyStr := buf.String()
		return fmt.Errorf("Pakk API Response Code: %s; Body: %s", resp.Status, bodyStr)
	}

	if err := json.NewDecoder(resp.Body).Decode(entity); err != nil {
		return err
	}

	return nil
}

type ListSpec struct {
	Filters          []Filter    `json:"filters"`
	FilterSets       []FilterSet `json:"filterSets"`
	Order            string      `json:"order"`
	OrderDescending  bool        `json:"orderDescending"`
	Limit            int         `json:"limit"`
	Offset           int         `json:"offset"`
	IncludeInactives bool        `json:"includeInactives"`
}

type FilterSet struct {
	Filters         []Filter `json:"filters"`
	FilterCombineOr bool     `json:"filterCombineOr"`
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
