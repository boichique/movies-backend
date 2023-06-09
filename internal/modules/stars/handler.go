package stars

import (
	"net/http"

	"github.com/boichique/movie-reviews/contracts"
	"github.com/boichique/movie-reviews/internal/config"
	"github.com/boichique/movie-reviews/internal/echox"
	"github.com/boichique/movie-reviews/internal/pagination"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service          *Service
	paginationConfig config.PaginationConfig
}

func NewHandler(service *Service, paginationConfig config.PaginationConfig) *Handler {
	return &Handler{
		service:          service,
		paginationConfig: paginationConfig,
	}
}

func (h *Handler) Create(c echo.Context) error {
	req, err := echox.BindAndValidate[contracts.CreateStarRequest](c)
	if err != nil {
		return err
	}

	star := &Star{
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		BirthDate:  req.BirthDate,
		BirthPlace: req.BirthPlace,
		DeathDate:  req.DeathDate,
		Bio:        req.Bio,
	}

	err = h.service.Create(c.Request().Context(), star)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, star)
}

func (h *Handler) GetStarsPaginated(c echo.Context) error {
	req, err := echox.BindAndValidate[contracts.GetStarsPaginatedRequest](c)
	if err != nil {
		return err
	}

	pagination.SetDefaults(&req.PaginatedRequest, h.paginationConfig)
	offset, limit := pagination.OffsetLimit(&req.PaginatedRequest)

	stars, total, err := h.service.GetStarsPaginated(c.Request().Context(), offset, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pagination.Response(&req.PaginatedRequest, total, stars))
}

func (h *Handler) GetByID(c echo.Context) error {
	req, err := echox.BindAndValidate[contracts.GetOrDeleteStarRequest](c)
	if err != nil {
		return err
	}

	star, err := h.service.GetByID(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, star)
}

func (h *Handler) Update(c echo.Context) error {
	req, err := echox.BindAndValidate[contracts.UpdateStarRequest](c)
	if err != nil {
		return err
	}
	star := &Star{
		ID:         req.ID,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		BirthDate:  req.BirthDate,
		BirthPlace: req.BirthPlace,
		DeathDate:  req.DeathDate,
		Bio:        req.Bio,
	}

	if err = h.service.Update(c.Request().Context(), star); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) Delete(c echo.Context) error {
	req, err := echox.BindAndValidate[contracts.GetOrDeleteStarRequest](c)
	if err != nil {
		return err
	}

	if err = h.service.Delete(c.Request().Context(), req.ID); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
