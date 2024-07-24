// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"tarkib.uz/internal/entity"
)

// //go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test
type (
	Auth interface {
		Login(ctx context.Context, admin *entity.Admin) (*entity.AdminLoginResponse, error)
		SuperAdminLogin(ctx context.Context, admin *entity.SuperAdmin) (*entity.SuperAdminLoginResponse, error)
		CreateAdmin(ctx context.Context, admin *entity.Admin) error
		DeleteAdmin(ctx context.Context, id string) error
		GetAllAdmins(ctx context.Context) ([]entity.Admin, error)
		EditAdmin(ctx context.Context, admin *entity.Admin) error
		GetAdminById(ctx context.Context, id string) (*entity.Admin, error)
		ChangeSuperAdminData(ctx context.Context, superAdmin *entity.SuperAdmin) error
	}

	AuthRepo interface {
		GetAdminData(ctx context.Context, Username string) (*entity.Admin, error)
		GetSuperAdminData(ctx context.Context, PhoneNumber string) (*entity.SuperAdmin, error)
		CreateAdmin(ctx context.Context, admin *entity.Admin) error
		DeleteAdmin(ctx context.Context, id string) error
		GetAllAdmins(ctx context.Context) ([]entity.Admin, error)
		EditAdmin(ctx context.Context, admin *entity.Admin) error
		GetAdminById(ctx context.Context, id string) (*entity.Admin, error)
		ChangeSuperAdminData(ctx context.Context, superAdmin *entity.SuperAdmin) error
	}

	News interface {
		CreateNews(ctx context.Context, request *entity.News) error
		GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error)
		DeleteNews(ctx context.Context, id string) error
		GetFilteredNews(ctx context.Context, request *entity.GetFilteredNewsRequest) ([]entity.News, error)
	}

	NewsRepo interface {
		CreateNews(ctx context.Context, request *entity.News) error
		GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error)
		DeleteNews(ctx context.Context, id string) error
		GetFilteredNews(ctx context.Context, request *entity.GetFilteredNewsRequest) ([]entity.News, error)
	}

	Category interface {
		GetAllCategories(ctx context.Context) ([]entity.Category, error)
	}

	CategoryRepo interface {
		GetAllCategories(ctx context.Context) ([]entity.Category, error)
	}
)
