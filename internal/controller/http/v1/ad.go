package v1

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"tarkib.uz/internal/entity"
	"tarkib.uz/internal/usecase"
	"tarkib.uz/pkg/logger"
	tokens "tarkib.uz/pkg/token"
)

type adRoutes struct {
	t usecase.AdUseCase
	l logger.Interface
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func newAdRoutes(handler *gin.RouterGroup, t usecase.AdUseCase, l logger.Interface) {
	r := &adRoutes{t, l}
	h := handler.Group("/ads")
	{
		h.POST("/", r.createAd)
		h.DELETE("/:id", r.deleteAd)
		h.PUT("/:id", r.updateAd)
		h.GET("/", r.getAd)
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
// @Security    BearerAuth
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
// @Security    BearerAuth
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

// @Summary		Update Ad
// @Description Edits ad by ID
// @Tags        ads
// @Produce     json
// @Success     204
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /ads/{id} [put]
func (r *adRoutes) updateAd(c *gin.Context) {
	var ad entity.Ad

	if err := c.ShouldBindJSON(&ad); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "Request body not matched", false)
		return
	}
	if err := r.t.UpdateAd(c.Request.Context(), &ad); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to delete ad", false)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary		Gets ad details
// @Description returns ads
// @Tags        ads
// @Produce     json
// @Success     200
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /ads [get]
func (r *adRoutes) getAd(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	fmt.Println(tokenStr)
	if tokenStr == "" {
		ad, err := r.t.GetAd(c.Request.Context(), &entity.GetAdRequest{
			IsAdmin: false,
		})

		if err != nil {
			r.l.Error(err)
			errorResponse(c, http.StatusInternalServerError, "Failed to get ad"+err.Error(), false)
			return
		}

		c.JSON(http.StatusOK, ad)
	} else {

		jwt := tokens.JWTHandler{
			SigninKey: "dfhdghkglioe",
			Token:     tokenStr,
		}

		claims, err := jwt.ExtractClaims()
		if err != nil {
			r.l.Error(err)
			errorResponse(c, http.StatusInternalServerError, "Failed to get aaad"+err.Error(), false)
			return
		}

		if claims["role"] == "super-admin" {
			ad, err := r.t.GetAd(c.Request.Context(), &entity.GetAdRequest{
				IsAdmin: true,
			})
			if err != nil {
				r.l.Error(err)
				errorResponse(c, http.StatusInternalServerError, "Failed to get ad", false)
				return
			}

			c.JSON(http.StatusOK, ad)
		}
	}

}
