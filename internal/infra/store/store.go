package store

import (
	"mocks/internal/core/model"
	"mocks/internal/infra/store/memory"
	"mocks/internal/infra/store/redis"
)

type Store interface {
	GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error)
	SetFoobar(req *model.FoobarRequest, resp *model.FoobarResponse) error
}

func NewStore(conf Config) (Store, error) {
	if conf.Memory != nil {
		return memory.NewStore(conf.Memory), nil
	}
	if conf.Redis != nil {
		return redis.NewStore(conf.Redis), nil
	}
	return nil, nil
}
