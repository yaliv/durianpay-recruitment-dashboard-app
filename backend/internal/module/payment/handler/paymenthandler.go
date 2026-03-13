package ph

import (
	"encoding/json"
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	pu "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
)

type PaymentHandler struct {
	paymentUC pu.PaymentUsecase
}

func NewPaymentHandler(paymentUC pu.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUC: paymentUC}
}

func (h *PaymentHandler) GetPaymentList(w http.ResponseWriter, r *http.Request, body openapigen.GetPaymentListParams) {
	payments, err := h.paymentUC.GetPaymentList(&body)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	paymentsRes := make([]openapigen.Payment, len(payments))

	for i, p := range payments {
		paymentsRes[i] = openapigen.Payment(*p)
	}

	err = json.NewEncoder(w).Encode(openapigen.PaymentListResponse{Payments: &paymentsRes})
	if err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}

func (h *PaymentHandler) GetPaymentSummary(w http.ResponseWriter, r *http.Request) {
	paymentSummaryRes, err := h.paymentUC.GetPaymentSummary()
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(paymentSummaryRes)
	if err != nil {
		transport.WriteAppError(w, entity.ErrorInternal("internal server error"))
		return
	}
}
