package service

import "github.com/injet-zhou/just-img-go-server/internal/dao"

type Service struct {
	Dao *dao.Dao
}

func NewService(dao *dao.Dao) *Service {
	return &Service{
		Dao: dao,
	}
}

func Default() *Service {
	return &Service{
		Dao: dao.Default(),
	}
}
