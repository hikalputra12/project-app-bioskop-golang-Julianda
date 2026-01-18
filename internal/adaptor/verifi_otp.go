package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type VerifyAdaptor struct {
	sessionRepo   usecase.SessionUsecaseInterface
	verifyUsecase usecase.VerifyUsecaseInterface
	log           *zap.Logger
}

func NewVerifyAdaptor(sessionRepo usecase.SessionUsecaseInterface, verifyUsecase usecase.VerifyUsecaseInterface, log *zap.Logger) *VerifyAdaptor {
	return &VerifyAdaptor{
		sessionRepo:   sessionRepo,
		verifyUsecase: verifyUsecase,
		log:           log,
	}
}

func (a *VerifyAdaptor) SomeVerifyFunction(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyOTP

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "unvalid format", nil)
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

	ctx := r.Context()
	sessionID, err := a.verifyUsecase.VerifyOTP(ctx, req)
	if err != nil {
		a.log.Error("failed login user on usecase",
			zap.Error(err),
		)
		utils.ResponseError(w, http.StatusUnauthorized, "Kode OTP salah atau sudah kadaluarsa", nil)
		return
	}

	//pembuatan cookie saat login

	expiryTime := 24 * 60 * 60
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   expiryTime,
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Login successful",
	})

}
