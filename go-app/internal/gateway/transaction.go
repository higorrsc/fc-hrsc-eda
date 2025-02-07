package gateway

import "github.com/higorrsc/fc-hrsc-eda/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
