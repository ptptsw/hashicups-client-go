package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllFood - Returns all user's food
func (c *Client) GetAllFood(authToken *string) (*[]Food, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	food := []Food{}
	err = json.Unmarshal(body, &food)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

// GetFood - Returns a specifc food
func (c *Client) GetFood(foodID string, authToken *string) (*Food, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/%s", c.HostURL, foodID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	food := Food{}
	err = json.Unmarshal(body, &food)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

// CreateFood - Create new food
func (c *Client) CreateFood(foodItems []FoodItem, authToken *string) (*Food, error) {
	rb, err := json.Marshal(foodItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/food", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	food := Food{}
	err = json.Unmarshal(body, &food)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

// UpdateFood - Updates an food
func (c *Client) UpdateFood(foodID string, foodItems []FoodItem, authToken *string) (*Food, error) {
	rb, err := json.Marshal(foodItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/food/%s", c.HostURL, foodID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	food := Food{}
	err = json.Unmarshal(body, &food)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

// DeleteFood - Deletes an food
func (c *Client) DeleteFood(foodID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/food/%s", c.HostURL, foodID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Deleted food" {
		return errors.New(string(body))
	}

	return nil
}
