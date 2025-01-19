package gateway

import "github.com/higorrsc/fc-hrsc-microservices/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
