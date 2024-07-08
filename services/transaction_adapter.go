package services

import (
	"Transaction/entities"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type transactionService struct {
	baseURL string
}

func NewTransactionService(baseURL string) TransactionService {
	return &transactionService{baseURL: baseURL}
}

func (s *transactionService) DistributeTransaction(payload entities.PayLoad) (string, error) {
	url := s.baseURL + "/broadcast"
	// TODO:convert payload to JSON
	jsonData, err := json.Marshal(&payload)
	if err != nil {
		return "", err
	}

	// TODO: POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//TODO: convert json to struct
	var serverRes = entities.ServerRespons{}
	if err := json.Unmarshal(body, &serverRes); err != nil {
		return "", err
	}

	return serverRes.Tx_Hash, nil
}

func (s *transactionService) CheckTransactionStatus(hash string) (string, error) {
	// TODO: GET
	url := s.baseURL + "/check/" + hash
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//TODO: Map json
	status := entities.ServerResponsStatus{}
	if err := json.Unmarshal(body, &status); err != nil {
		return "", err
	}
	return status.Tx_Status, nil
}
