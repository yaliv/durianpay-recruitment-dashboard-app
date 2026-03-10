package pu

import (
	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	pr "github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type (
	PaymentUsecase interface {
		GetPaymentList(params *openapigen.GetDashboardV1PaymentsParams) ([]*entity.Payment, error)
	}

	Payment struct {
		repo pr.PaymentRepository
	}
)

func NewPaymentUsecase(repo pr.PaymentRepository) *Payment {
	return &Payment{repo: repo}
}

func (uc *Payment) GetPaymentList(params *openapigen.GetDashboardV1PaymentsParams) ([]*entity.Payment, error) {
	if params.Id != nil {
		p, err := uc.repo.GetPaymentByID(*params.Id)
		if err != nil {
			return nil, err
		}

		return []*entity.Payment{p}, nil
	}

	ps, err := uc.repo.AllPayments(params.Status, params.Sort)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
