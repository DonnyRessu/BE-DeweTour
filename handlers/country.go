package handlers

import (
	"net/http"
	"strconv"
	tripdto "week2/dto/country"
	dto "week2/dto/result"
	"week2/models"
	"week2/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCountry struct {
	CountryRepository repositories.CountryRepository
}

func HandlersCountry(CountryRepository repositories.CountryRepository) *handlerCountry {
	return &handlerCountry{CountryRepository}
}

func (h *handlerCountry) FindCountry(c echo.Context) error {
	users, err := h.CountryRepository.FindCountry()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *handlerCountry) GetCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: user})
}

func (h *handlerCountry) CreateCountry(c echo.Context) error {
	get := new(tripdto.CreateCountryRequest)
	if err := c.Bind(get); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	validation := validator.New()
	err := validation.Struct(get)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	user := models.Country{
		ID:   get.ID,
		Name: get.Name,
	}
	data, err := h.CountryRepository.CreateCountry(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCountry) UpdateCountry(c echo.Context) error {
	get := new(tripdto.UpdateCountryRequest)
	if err := c.Bind(get); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if get.ID != 0 {
		user.ID = get.ID
	}
	if get.Name != "" {
		user.Name = get.Name
	}

	data, err := h.CountryRepository.UpdateCountry(user, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCountry) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CountryRepository.DeleteCountry(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func convertResponseCat(u models.Country) tripdto.CountryResponse {
	return tripdto.CountryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
