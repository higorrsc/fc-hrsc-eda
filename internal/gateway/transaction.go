package gateway

import "github.com/higorrsc/fc-hrsc-microservices/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
