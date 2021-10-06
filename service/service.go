package service

import (
	"pipe-mbx/model"
	"pipe-mbx/repo"
)

type svc struct {
	opts Opts
	repo repo.SaveRepo
}

// Opts is an option for service
type Opts struct {
	DataType    string
	Outcome     model.OutcomeType
	RawDataPath string
	SavePath    string
}

type Service interface {
	Run() error
}

func NewService(r repo.SaveRepo, o Opts) (Service, error) {
	return &svc{
		opts: o,
		repo: r,
	}, nil
}
