package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"net/http"

	"go.uber.org/zap"
)

type PaymentMethodAdaptor struct {
	PaymentMethodUsecase usecase.PaymentMethodUsecaseInterface
	log                  *zap.Logger
}

func NewPaymentMethodAdaptor(paymentMethodUsecase usecase.PaymentMethodUsecaseInterface, log *zap.Logger) *PaymentMethodAdaptor {
	return &PaymentMethodAdaptor{
		PaymentMethodUsecase: paymentMethodUsecase,
		log:                  log,
	}
}

func (p *PaymentMethodAdaptor) GetAllPaymentMethods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paymentMethod, err := p.PaymentMethodUsecase.GetAllPaymentMethods(ctx)
	if err != nil {
		p.log.Error("failed get all payment methods on usecase")
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch cinemas: "+err.Error(), nil)
		return
	}
	var response []dto.PaymentMethodResponse
	for _, item := range paymentMethod {
		response = append(response, dto.PaymentMethodResponse{
			Name: item.MethodName,
			Logo: item.Logo,
		})

	}
	utils.ResponseSuccess(w, http.StatusOK, "Succses get data", response)
}
