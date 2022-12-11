package usecase

import (
	"lj/app/domain"
)

type TinyUsecase struct {
	domain *domain.Domain
}

func NewTinyUsecase(domain *domain.Domain) *TinyUsecase {
	return &TinyUsecase{domain: domain}
}
