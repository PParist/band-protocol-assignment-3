package services

import "Transaction/entities"

type TransactionService interface {
	DistributeTransaction(entities.PayLoad) (string, error)
	CheckTransactionStatus(hash string) (string, error)
}
