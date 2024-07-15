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
		h.GET("/getall/withcategory", r.getAllNewsWithCategoryNames)
	}
}

// @Summary     Create News
// @Description This method for creating a new news.
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
		Name:          body.Name,
		Description:   body.Description,
		ImageURL:      body.ImageURL,
		SubCategoryID: body.SubCategoryID,
	}); err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, "Kechirasiz, serverda muammolar bo'lyapti")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Yangilik muvaffaqiyatli yaratildi.",
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

// @Summary		Get All News By Category
// @Description This method retrieves all news with category and subcategory names.
// @ID          get-all-news-with-category
// @Tags  	    news
// @Accept      json
// @Produce     json
// @Param       page  query int true  "Page number"
// @Param       limit query int true  "Number of items per page"
// @Param       subcategory_id query int true  "Subcategory id for getting all news by subcategory and category"
// @Success     200 {object} []models.NewsWithCategoryNames
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /news/getall/withcategory [get]
func (n *newsRoutes) getAllNewsWithCategoryNames(c *gin.Context) {
	var (
		page          string
		limit         string
		subcategoryId string
	)

	page = c.Query("page")
	limit = c.Query("limit")
	subcategoryId = c.Query("subcategory_id")

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

	subcategoryIdInt, err := strconv.Atoi(subcategoryId)
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	news, err := n.t.GetAllNewsByCategory(c.Request.Context(), &entity.GetNewsBySubCategory{
		Page:          pageInt,
		Limit:         limitInt,
		SubCategoryId: subcategoryIdInt,
	})
	if err != nil {
		n.l.Error(err)
		errorResponse(c, http.StatusInternalServerError, models.ErrServerProblems)
		return
	}

	c.JSON(http.StatusOK, news)
}
