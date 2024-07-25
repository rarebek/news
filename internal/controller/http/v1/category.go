package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tarkib.uz/internal/controller/http/models"
	"tarkib.uz/internal/usecase"
	"tarkib.uz/pkg/logger"
)

type categoryRoutes struct {
	t usecase.CategoryUseCase
	l logger.Interface
}

func newCategoryRoutes(handler *gin.RouterGroup, t usecase.CategoryUseCase, l logger.Interface) {
	r := &categoryRoutes{t, l}

	h := handler.Group("/category")
	{
		h.GET("/getall", r.getAllCategories)
	}
}

// @Summary     Get All Categories
// @Description This method retrieves all categories with their subcategories.
// @ID          get-all-categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Success     200 {array} models.CategoryResponse
// @Failure     500 {object} response
// @Router      /category/getall [get]
func (n *categoryRoutes) getAllCategories(c *gin.Context) {
	categories, err := n.t.GetAllCategories(c.Request.Context())
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Kechirasiz, serverda muammolar bo'lyapti", false)
		return
	}

	var categoryResponses []models.CategoryResponse
	for _, category := range categories {
		var subCategoryResponses []models.SubCategoryResponse
		for _, subCategory := range category.SubCategories {
			subCategoryResponses = append(subCategoryResponses, models.SubCategoryResponse{
				ID:   subCategory.ID, // Use string for UUID
				Name: subCategory.Name,
			})
		}
		categoryResponses = append(categoryResponses, models.CategoryResponse{
			ID:            category.ID, // Use string for UUID
			Name:          category.Name,
			SubCategories: subCategoryResponses,
		})
	}

	c.JSON(http.StatusOK, categoryResponses)
}
