package start

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/config"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ar "github.com/durianpay/fullstack-boilerplate/internal/module/auth/repository"
	au "github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	ph "github.com/durianpay/fullstack-boilerplate/internal/module/payment/handler"
	pr "github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
	pu "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
)

func SetupAPIHandler(db *sqlx.DB) *api.APIHandler {
	JwtExpiredDuration, err := time.ParseDuration(config.JwtExpired)
	if err != nil {
		panic(err)
	}

	userRepo := ar.NewUserRepo(db)

	authUC := au.NewAuthUsecase(userRepo, config.JwtSecret, JwtExpiredDuration)

	authH := ah.NewAuthHandler(authUC)

	paymentRepo := pr.NewPaymentRepo(db)

	paymentUC := pu.NewPaymentUsecase(paymentRepo)

	paymentH := ph.NewPaymentHandler(paymentUC)

	return &api.APIHandler{
		Auth:    authH,
		Payment: paymentH,
	}
}
