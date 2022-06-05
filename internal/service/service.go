package service

import (
	"github.com/google/wire"
	v1 "kartos-realworld/api/realworld/v1"
	"kartos-realworld/internal/biz"

)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewRealworldService(uc *biz.GreeterUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}