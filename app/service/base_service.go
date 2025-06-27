package service

type BaseServiceInterface interface {
}

type BaseService struct {
}

func NewBaseService() BaseService {
	return BaseService{}
}
