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

func (h *PaymentHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request) {
	var req openapigen.GetDashboardV1PaymentsParams
	getParams(r, &req)

	payments, err := h.paymentUC.GetPaymentList(&req)
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

func getParams(r *http.Request, req *openapigen.GetDashboardV1PaymentsParams) {
	urlV := r.URL.Query()

	if urlV.Has("sort") {
		*req.Sort = openapigen.Sort(urlV.Get("sort"))
	}
	if urlV.Has("status") {
		*req.Status = urlV.Get("status")
	}
	if urlV.Has("id") {
		*req.Id = urlV.Get("id")
	}
}
