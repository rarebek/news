package v1

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

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
		h.GET("/admin/getall", r.getAllAdmins)
		h.POST("/upload", r.upload)
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
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusBadRequest, "Bunday admin topilmadi.")
		case "xato parol kiritdingiz":
			r.l.Warn(err.Error())
			errorResponse(c, http.StatusUnauthorized, "Username yoki parol xato kiritildi.")
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
// @Param       request body models.SuperAdminLoginRequest true "Phone Number and Password"
// @Success     200 {object} models.AdminLoginResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /auth/superadmin/login [post]
func (r *authRoutes) superAdminLogin(c *gin.Context) {
	var request models.SuperAdminLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.InvalidRequestBody)
		return
	}

	token, err := r.t.SuperAdminLogin(c.Request.Context(), &entity.SuperAdmin{
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
		Username: request.Username,
		Password: request.Password,
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
// @Success     200 {object} response
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /auth/admin/delete [delete]
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

// @Summary     Get All Admins
// @Description Gets All Admins
// @ID          get-all-admins
// @Tags  	    superadmin
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Admin
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Security    BearerAuth
// @Router      /auth/admin/getall [get]
func (r *authRoutes) getAllAdmins(c *gin.Context) {
	admins, err := r.t.GetAllAdmins(c.Request.Context())
	if err != nil {
		r.l.Error(err)
		errorResponse(c, http.StatusBadRequest, models.ErrServerProblems)
		return
	}

	c.JSON(http.StatusOK, admins)
}

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Summary       Image upload
// @Description   Api for image upload
// @Tags          file-upload
// @Accept        multipart/form-data
// @Produce       json
// @Param         file formData file true "Image"
// @Param         type formData string true "Bucket type to put image"
// @Success       200 {object} string
// @Failure       400 {object} string
// @Failure       500 {object} string
// @Router        /file/upload [post]
func (f *authRoutes) upload(c *gin.Context) {
	pp.Println("This method")
	// Parse form fields
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		f.l.Error(err, "http - v1 - fileupload - ParseMultipartForm")
		errorResponse(c, http.StatusBadRequest, "Failed to parse multipart form")
		return
	}

	// Extract type from form
	bucketType := c.Request.FormValue("type")
	if bucketType == "" {
		errorResponse(c, http.StatusBadRequest, "Bucket type is required")
		return
	}
	fmt.Println("a")
	// Validate file field and retrieve file data
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		f.l.Error(err, "http - v1 - fileupload - FormFile")
		errorResponse(c, http.StatusBadRequest, "Failed to retrieve file")
		return
	}
	defer file.Close()

	// Validate file extension
	ext := filepath.Ext(fileHeader.Filename)
	allowedExts := map[string]bool{".png": true, ".jpg": true, ".svg": true, ".jpeg": true, ".JPG": true, ".PNG": true}
	if !allowedExts[ext] {
		f.l.Error(errors.New("invalid file extension"), "http - v1 - fileupload - filepath.Ext")
		errorResponse(c, http.StatusBadRequest, "Only PNG, JPG, JPEG, and SVG files are allowed")
		return
	}

	// Prepare MinIO client and upload parameters
	endpoint := os.Getenv("SERVER_IP")
	accessKeyID := "nodirbek"
	secretAccessKey := "nodirbek"
	bucketName := bucketType
	minioClient, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		f.l.Error(err, "http - v1 - fileupload - minio.New")
		errorResponse(c, http.StatusInternalServerError, "Failed to initialize MinIO client")
		return
	}

	// Generate unique object name
	id := uuid.New().String()
	objectName := id + ext
	contentType := "image/jpeg"

	// Upload file to MinIO
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		f.l.Error(err, "http - v1 - fileupload - minioClient.PutObject")
		errorResponse(c, http.StatusInternalServerError, "Failed to upload file to MinIO")
		return
	}

	// Construct MinIO URL
	minioURL := fmt.Sprintf("https://%s/%s/%s", endpoint, bucketName, objectName)

	// Respond with success message containing URL
	c.JSON(http.StatusOK, gin.H{
		"url": minioURL,
	})
}
