package gateway

import "github.com/higorrsc/fc-hrsc-eda/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
