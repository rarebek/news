package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"tarkib.uz/internal/entity"
	"tarkib.uz/internal/usecase"
	"tarkib.uz/pkg/logger"
)

type adRoutes struct {
	t usecase.AdUseCase
	l logger.Interface
}

func newAdRoutes(handler *gin.RouterGroup, t usecase.AdUseCase, l logger.Interface) {
	r := &adRoutes{t, l}
	h := handler.Group("/ads")
	{
		h.POST("/", r.createAd)
		h.DELETE("/:id", r.deleteAd)
	}
}

// @Summary     Create a new ad
// @Description Create a new ad with the given details
// @Tags        ads
// @Accept      json
// @Produce     json
// @Param       ad body entity.CreateAdRequest true "Ad details"
// @Success     201 {object} entity.Ad
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /ads [post]
func (r *adRoutes) createAd(c *gin.Context) {
	var ad entity.CreateAdRequest
	if err := c.ShouldBindJSON(&ad); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "Invalid request body", false)
		return
	}
	id := uuid.NewString()
	if err := r.t.CreateAd(c.Request.Context(), &entity.Ad{
		ID:          id,
		Title:       ad.Title,
		Description: ad.Description,
		ImageURL:    ad.ImageURL,
	}); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create ad", false)
		return
	}

	var response entity.Ad
	response.ID = id
	response.Description = ad.Description
	response.ImageURL = ad.ImageURL
	response.Title = ad.Title

	c.JSON(http.StatusCreated, response)
}

// @Summary     Delete an ad
// @Description Delete an ad by ID
// @Tags        ads
// @Produce     json
// @Param       id path string true "Ad ID"
// @Success     204
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /ads/{id} [delete]
func (r *adRoutes) deleteAd(c *gin.Context) {
	id := c.Param("id")

	if err := r.t.DeleteAd(c.Request.Context(), id); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to delete ad", false)
		return
	}

	c.Status(http.StatusNoContent)
}
