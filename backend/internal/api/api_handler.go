package api

import (
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ph "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth    *ah.AuthHandler
	Payment *ph.PaymentHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.Auth.Login(w, r)
}

func (h *APIHandler) GetPaymentList(w http.ResponseWriter, r *http.Request, body openapigen.GetPaymentListParams) {
	h.Payment.GetPaymentList(w, r, body)
}

func (h *APIHandler) GetPaymentSummary(w http.ResponseWriter, r *http.Request) {
	h.Payment.GetPaymentSummary(w, r)
}
