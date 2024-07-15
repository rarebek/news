// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"tarkib.uz/internal/entity"
)

// //go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test
type (
	Auth interface {
		Login(ctx context.Context, admin *entity.Admin) (string, error)
		SuperAdminLogin(ctx context.Context, admin *entity.Admin) (string, error)
		CreateAdmin(ctx context.Context, admin *entity.Admin) error
	}

	AuthRepo interface {
		GetAdminData(ctx context.Context, PhoneNumber string) (*entity.Admin, error)
		GetSuperAdminData(ctx context.Context, PhoneNumber string) (*entity.Admin, error)
		CreateAdmin(ctx context.Context, admin *entity.Admin) error
	}

	News interface {
		CreateNews(ctx context.Context, request *entity.News) error
		GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error)
		GetAllNewsByCategory(ctx context.Context, request *entity.GetNewsBySubCategory) ([]entity.NewsWithCategoryNames, error)
	}

	NewsRepo interface {
		CreateNews(ctx context.Context, request *entity.News) error
		GetAllNews(ctx context.Context, request *entity.GetAllNewsRequest) ([]entity.News, error)
		GetAllNewsByCategory(ctx context.Context, request *entity.GetNewsBySubCategory) ([]entity.NewsWithCategoryNames, error)
	}

	Category interface {
		GetAllCategories(ctx context.Context) ([]entity.Category, error)
	}

	CategoryRepo interface {
		GetAllCategories(ctx context.Context) ([]entity.Category, error)
	}
)
