package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"tarkib.uz/internal/controller/http/models"
	"tarkib.uz/internal/entity"
	"tarkib.uz/internal/usecase"
	"tarkib.uz/pkg/logger"
)

type newsRoutes struct {
	t usecase.NewsUseCase
	l logger.Interface
}

func newNewsRoutes(handler *gin.RouterGroup, t usecase.NewsUseCase, l logger.Interface) {
	r := &newsRoutes{t, l}

	h := handler.Group("/news")
	{
		h.POST("/create", r.create)
		h.GET("/getall", r.getAllNews)
		h.DELETE("/delete/:id", r.deleteNews)
	}
}

// @Summary     Create News
// @Description This method for creating a news
// @ID          create-news
// @Tags  	    news
// @Accept      json
// @Produce     json
// @Param       request body models.News true "News details"
// @Success     200 {object} models.Message
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /news/create [post]
func (n *newsRoutes) create(c *gin.Context) {
	var body models.News
	if err := c.ShouldBindJSON(&body); err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := n.t.CreateNews(c.Request.Context(), &entity.News{
		Name:           body.Name,
		Description:    body.Description,
		ImageURL:       body.ImageURL,
		SubCategoryIDs: body.SubCategoryIDs,
	}); err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Kechirasiz, serverda muammolar bo'lyapti")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Yangilik muvaffaqiyatli yaratildi.",
	})
}

// @Summary     Delete News
// @Description This method deleting news
// @ID          delete-news
// @Tags  	    news
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of the news to delete"
// @Success     200 {object} models.Message
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /news/delete/{id} [delete]
func (n *newsRoutes) deleteNews(c *gin.Context) {
	id := c.Param("id")
	if err := n.t.DeleteNews(c.Request.Context(), id); err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Kechirasiz, serverda muammolar bo'lyapti")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Yangilik muvaffaqiyatli o'chirildi.",
	})
}

// @Summary		Get All News
// @Description This method retrieves all news with pagination.
// @ID          get-all-news
// @Tags  	    news
// @Accept      json
// @Produce     json
// @Param       page  query int true  "Page number"
// @Param       limit query int true  "Number of items per page"
// @Success     200 {object} []models.News
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /news/getall [get]
func (n *newsRoutes) getAllNews(c *gin.Context) {
	var (
		page  string
		limit string
	)

	page = c.Query("page")
	limit = c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	news, err := n.t.GetAllNews(c.Request.Context(), &entity.GetAllNewsRequest{
		Page:  pageInt,
		Limit: limitInt,
	})
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, models.ErrServerProblems)
		return
	}

	c.JSON(http.StatusOK, news)
}

// @Summary		Get Filtered News
// @Description This method retrieves news based on optional filters (subcategory IDs and category ID) with pagination.
// @ID          get-filtered-news
// @Tags  	    news
// @Accept      json
// @Produce     json
// @Param       sub_category_ids query []string false "List of subcategory IDs"
// @Param       category_id     query string   false "Category ID"
// @Param       page            query int      true  "Page number"
// @Param       limit           query int      true  "Number of items per page"
// @Success     200 {object} []models.News
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /news/filtered [get]
func (n *newsRoutes) getFilteredNews(c *gin.Context) {
	subCategoryIDs := c.QueryArray("sub_category_ids")
	categoryID := c.Query("category_id")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "Invalid page number")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, "Invalid limit number")
		return
	}

	news, err := n.t.GetFilteredNews(c.Request.Context(), &entity.GetFilteredNewsRequest{
		SubCategoryIDs: subCategoryIDs,
		CategoryID:     categoryID,
		Page:           page,
		Limit:          limit,
	})
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Kechirasiz, serverda muammolar bo'lyapti")
		return
	}

	c.JSON(http.StatusOK, news)
}
