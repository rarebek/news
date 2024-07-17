package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tarkib.uz/internal/controller/http/models"
	"tarkib.uz/internal/entity"
	"tarkib.uz/internal/usecase"
	"tarkib.uz/pkg/logger"
)

type authRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
	r := &authRoutes{t, l}

	h := handler.Group("/auth")
	{
		h.POST("/admin/login", r.login)
		h.POST("/superadmin/login", r.superAdminLogin)
		h.POST("/admin/create", r.createAdmin)
		h.DELETE("/admin/delete/:id", r.deleteAdmin)
	}
}

// @Summary     Login
// @Description Authenticates an admin and returns an access token on successful login.
// @ID          admin-login
// @Tags  	    admin
// @Accept      json
// @Produce     json
// @Param       request body models.AdminLoginRequest true "Phone Number and Password"
// @Success     200 {object} models.AdminLoginResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /auth/admin/login [post]
func (r *authRoutes) login(c *gin.Context) {
	var request models.AdminLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.InvalidRequestBody)
		return
	}

	token, err := r.t.Login(c.Request.Context(), &entity.Admin{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	})

	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusBadRequest, "Bunday admin topilmadi.")
		case "xato parol kiritdingiz":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusUnauthorized, "Telefon raqam yoki parol xato kiritildi.")
		default:
			r.l.Error(err)
			errorResponse(c, http.StatusInternalServerError, models.ErrServerProblems)
		}
		return
	}

	c.JSON(http.StatusOK, models.AdminLoginResponse{
		AccessToken: token,
	})
}

// @Summary     Super Admin Login
// @Description Authenticates a super admin and returns an access token on successful login.
// @ID          superadmin-login
// @Tags  	    superadmin
// @Accept      json
// @Produce     json
// @Param       request body models.AdminLoginRequest true "Phone Number and Password"
// @Success     200 {object} models.AdminLoginResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /auth/superadmin/login [post]
func (r *authRoutes) superAdminLogin(c *gin.Context) {
	var request models.AdminLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.InvalidRequestBody)
		return
	}

	token, err := r.t.SuperAdminLogin(c.Request.Context(), &entity.Admin{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	})

	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusBadRequest, "Bunday admin topilmadi.")
		case "xato parol kiritdingiz":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusUnauthorized, "Telefon raqam yoki parol xato kiritildi.")
		default:
			r.l.Error(err)
			errorResponse(c, http.StatusInternalServerError, models.ErrServerProblems)
		}
		return
	}

	c.JSON(http.StatusOK, models.AdminLoginResponse{
		AccessToken: token,
	})
}

// @Summary     Create Admin
// @Description Creates an admin
// @ID          superadmin-create-admin
// @Tags  	    superadmin
// @Accept      json
// @Produce     json
// @Param       request body models.AdminLoginRequest true "Phone Number and Password to create Admin"
// @Success     200 {object} models.AdminLoginResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /auth/admin/create [post]
func (r *authRoutes) createAdmin(c *gin.Context) {
	var request models.AdminLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.InvalidRequestBody)
		return
	}

	if err := r.t.CreateAdmin(c.Request.Context(), &entity.Admin{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin muvaffaqiyatli yaratildi",
	})
}

// @Summary     Delete Admin
// @Description This method deletes admin.
// @ID          superadmin-delete-admin
// @Tags  	    superadmin
// @Accept      json
// @Produce     json
// @Param       id path int true "ID of the admin to delete"
// @Success     200 {object} models.AdminLoginResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /auth/admin/delete [post]
func (r *authRoutes) deleteAdmin(c *gin.Context) {
	id := c.Param("id")

	if err := r.t.DeleteAdmin(c.Request.Context(), id); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin muvaffaqiyatli o'chirildi",
	})
}
