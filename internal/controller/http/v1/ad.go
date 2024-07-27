package v1

import (
	"net/http"
	"time"

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
		h.DELETE("/expired", r.deleteExpiredAds)
	}
}

// @Summary     Create a new ad
// @Description Create a new ad with the given details
// @Tags        ads
// @Accept      json
// @Produce     json
// @Param       ad body entity.Ad true "Ad details"
// @Success     201 {object} entity.Ad
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /ads [post]
func (r *adRoutes) createAd(c *gin.Context) {
	var ad entity.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "Invalid request body", false)
		return
	}

	ad.ID = uuid.NewString()
	ad.ExpirationTime = time.Now().Add(getDuration(ad.Duration))

	if err := r.t.CreateAd(c.Request.Context(), &ad); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create ad", false)
		return
	}

	c.JSON(http.StatusCreated, ad)
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

// @Summary     Delete expired ads
// @Description Delete all ads that have expired
// @Tags        ads
// @Produce     json
// @Success     204
// @Failure     500 {object} response
// @Router      /ads/expired [delete]
func (r *adRoutes) deleteExpiredAds(c *gin.Context) {
	if err := r.t.DeleteExpiredAds(c.Request.Context()); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to delete expired ads", false)
		return
	}

	c.Status(http.StatusNoContent)
}

func getDuration(option string) time.Duration {
	switch option {
	case "1 day":
		return 24 * time.Hour
	case "2 days":
		return 48 * time.Hour
	case "3 days":
		return 72 * time.Hour
	case "1 week":
		return 7 * 24 * time.Hour
	case "2 weeks":
		return 14 * 24 * time.Hour
	case "monthly":
		return 30 * 24 * time.Hour
	default:
		return 0
	}
}
