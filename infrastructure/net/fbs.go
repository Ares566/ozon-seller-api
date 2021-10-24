package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"ozon-seller-api/domain/entity"
)

type OzonClient struct {
	ClientID string
	Key      string
	Client   *http.Client
}

const OzoneHost = "https://api-seller.ozon.ru"

func NewFBSClient(clientID string, key string) (*OzonClient, error) {
	if clientID == "" || key == "" {
		argErros := errors.New("нужно указать верные clientID и key")
		return nil, argErros
	}

	return &OzonClient{
		ClientID: clientID,
		Key:      key,
		Client:   &http.Client{},
	}, nil

}

func (c *OzonClient) SendBalance(sbr *entity.StockBalanceRequest) (entity.StockBalanceResponse, error) {
	// https://api-seller.ozon.ru/v2/products/stocks

	req, err := c.newPost("/v2/products/stocks", sbr)

	var result entity.StockBalanceResponse
	err = c.doRequest(req, &result)

	return result, err

}
func (c *OzonClient) GetUnfulfilledList() ([]entity.Posting, error) {
	//https://api-seller.ozon.ru/v3/posting/fbs/unfulfilled/list

	const (
		pageSize  = 100
		startPage = 1
	)

	totalElements := -1

	var list []entity.Posting

	for page := startPage; len(list) != totalElements; page++ {
		req, err := c.newPost("/v3/posting/fbs/unfulfilled/list", entity.UFFLRequest{
			Limit:  pageSize,
			Offset: (page - 1) * pageSize,
		})
		if err != nil {
			return list, err
		}

		var respStruct entity.UFFLResponse

		err = c.doRequest(req, &respStruct)
		if err != nil {
			return list, err
		}

		totalElements = respStruct.Result.Count
		list = append(list, respStruct.Result.Postings...)

	}

	return list, nil
}

func (c *OzonClient) newPost(path string, v interface{}) (*http.Request, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, OzoneHost+path, bytes.NewBuffer(b))
	if err != nil {
		return req, err
	}
	req.Header.Set("Client-Id", c.ClientID)
	req.Header.Set("Api-Key", c.Key)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req, nil
}

func (c *OzonClient) doRequest(req *http.Request, v interface{}) error {

	resp, err := c.Client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
