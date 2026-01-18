package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CinemaAdaptor struct {
	CinemaUsecase usecase.CinemaUsecaseInterface
	log           *zap.Logger
}

func NewCinemaAdaptor(CinemaUsecase usecase.CinemaUsecaseInterface, log *zap.Logger) *CinemaAdaptor {
	return &CinemaAdaptor{
		CinemaUsecase: CinemaUsecase,
		log:           log,
	}
}

// get all cinemas
func (c *CinemaAdaptor) GetAllCinemas(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
		return
	}

	// config limit pagination
	limit := 5
	ctx := r.Context()
	// Get data Cinemas form usecase all cinemas
	cinemas, pagination, err := c.CinemaUsecase.GetAllCinemas(ctx, page, limit)
	if err != nil {
		c.log.Error("failed get all cinemas on usecase")
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch cinemas: "+err.Error(), nil)
		return
	}
	var response []dto.CinemaResponse
	for _, item := range cinemas {
		response = append(response, dto.CinemaResponse{
			Name:     item.Name,
			Location: item.Location,
		})

	}
	utils.ResponsePagination(w, http.StatusOK, "success get data", response, *pagination)

}

// get cinemas by id
func (c *CinemaAdaptor) GetcinemasById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "cinemaId")
	cinemaID, _ := strconv.Atoi(id)

	// Get data cinemass form service all cinemass
	ctx := r.Context()
	cinemas, err := c.CinemaUsecase.GetCinemaByID(ctx, cinemaID)

	if err != nil {
		c.log.Error("failed get cinemas by id on usecase",
			zap.Error(err),
		)
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch assignments: "+err.Error(), nil)
		return
	}

	var response dto.CinemaResponse
	response = dto.CinemaResponse{
		Name:     cinemas.Name,
		Location: cinemas.Location,
	}
	utils.ResponseSuccess(w, http.StatusOK, "success get data", response)
}

// get seat cinema by date and time
func (c *CinemaAdaptor) GetSeatCinema(w http.ResponseWriter, r *http.Request) {
	cinemID, err := strconv.Atoi(chi.URLParam(r, "cinemaId"))
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid cinema ID", nil)
		return
	}
	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")
	if date == "" || time == "" {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Date and time are required", nil)
		return
	}

	ctx := r.Context()
	// Get data Cinemas form usecase all cinemas
	seats, err := c.CinemaUsecase.GetSeatCinema(ctx, cinemID, date, time)
	if err != nil {
		c.log.Error("failed gewt all cinemas on usecase")
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch cinemas: "+err.Error(), nil)
		return
	}
	var response []dto.SeatResponse
	for _, item := range seats {
		response = append(response, dto.SeatResponse{
			SeatNumber: item.SeatNumber,
			IsAvaiable: item.IsAvaiable,
		})

	}
	utils.ResponseSuccess(w, http.StatusOK, "Succses get data", response)

}
