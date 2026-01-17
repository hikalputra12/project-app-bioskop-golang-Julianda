package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type AuthAdaptor struct {
	AuthUsecase    usecase.AuthUsecaseInterface
	SessionUsecase usecase.SessionUsecaseInterface
	log            *zap.Logger
}

func NewAuthAdaptor(uc usecase.AuthUsecaseInterface, SessionUsecase usecase.SessionUsecaseInterface, log *zap.Logger) *AuthAdaptor {
	return &AuthAdaptor{
		AuthUsecase:    uc,
		SessionUsecase: SessionUsecase,
		log:            log,
	}
}

func (a *AuthAdaptor) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

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
	sessionID, err := a.AuthUsecase.Login(ctx, req)
	if err != nil {
		a.log.Error("failed login user on usecase",
			zap.Error(err),
		)
		utils.ResponseError(w, http.StatusUnauthorized, "email or password incorrect", nil)
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

func (a *AuthAdaptor) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	ctx := r.Context()
	if err != nil {
		// Jika cookie tidak ada, anggap saja user sudah logout
		utils.ResponseError(w, http.StatusUnauthorized, "User tidak terautentikasi", nil)
		return
	}
	//pembuatan revoke saat logout
	revokedSession := cookie.Value

	err = a.SessionUsecase.RevokedSession(ctx, revokedSession)
	if err != nil {
		a.log.Error("failed revoke session on usercase",
			zap.Error(err),
		)
		utils.ResponseError(w, http.StatusInternalServerError, "failed to proccess logout session", nil)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Logout successful",
	})

}
