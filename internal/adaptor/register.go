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
		// Pastikan error message dari usecase (seperti "email sudah terdaftar") diteruskan jika perlu
		utils.ResponseError(w, http.StatusBadRequest, "Gagal registrasi: "+err.Error(), nil)
		return
	}

	// --- BAGIAN INI YANG DIUBAH ---
	w.Header().Set("Content-Type", "application/json")

	// 1. Ganti jadi 201 Created
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		// 2. Pesan yang memandu User
		"message": "Registrasi berhasil. Kode OTP telah dikirim ke email Anda, silakan verifikasi.",
		// 3. (Opsional) Data tambahan buat Frontend
		"data": map[string]string{
			"email": req.Email,
		},
	})

	h.log.Info("sukses membuat user baru dan memicu pengiriman OTP", zap.String("email", req.Email))
}
