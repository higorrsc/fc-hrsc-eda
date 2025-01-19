package gateway

import "github.com/higorrsc/fc-hrsc-microservices/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
