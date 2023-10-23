package gateway

import "github.com/joferreira/fc-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(Client *entity.Client) error
}
