package handler

import (
	"Transaction/entities"
	"Transaction/services"
)

type transactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) transactionHandler {
	return transactionHandler{service: service}
}

func (h *transactionHandler) BroadcastTransaction(payload entities.PayLoad) (string, error) {
	if result, err := h.service.DistributeTransaction(payload); err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func (h *transactionHandler) TransactionStatusMonitoring(hash string) (string, error) {

	if result, err := h.service.CheckTransactionStatus(hash); err != nil {
		return "", err
	} else {
		return result, nil
	}
}
