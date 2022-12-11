package usecase

import (
	"lj/app/domain"
)

type Usecase struct {
	domain *domain.Domain
}

func NewUsecase(domain *domain.Domain) *Usecase {
	return &Usecase{domain: domain}
}
