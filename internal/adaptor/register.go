package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type RegisterAdaptor struct {
	RegisterUseCase usecase.RegisterUseCaseInterface
	log             *zap.Logger
}

func NewRegisterAdaptor(registerUseCase usecase.RegisterUseCaseInterface, log *zap.Logger) *RegisterAdaptor {
	return &RegisterAdaptor{
		RegisterUseCase: registerUseCase,
		log:             log,
	}
}

func (h *RegisterAdaptor) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	ctx := r.Context()

	//mengubah json body ke struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}
	validationErrors, err := utils.ValidateErrors(req)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Validation failed",
			"errors":  validationErrors,
		})
		return
	}

	err = h.RegisterUseCase.RegisterAccount(ctx, req)
	if err != nil {
		h.log.Error("failed register user on service",
			zap.Error(err),
		)
		utils.ResponseError(w, http.StatusBadRequest, "input tidak sesuai format yang di tentukan", nil)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "register succesfully",
	})
	h.log.Info("sukses membuat user baru")
}
