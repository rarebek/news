package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"tarkib.uz/config"
	"tarkib.uz/internal/entity"
	tokens "tarkib.uz/pkg/token"
)

type AuthUseCase struct {
	repo        AuthRepo
	cfg         *config.Config
	RedisClient *redis.Client
}

func NewAuthUseCase(r AuthRepo, cfg *config.Config) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
		cfg:  cfg,
	}
}

func (uc *AuthUseCase) Login(ctx context.Context, request *entity.Admin) (string, error) {
	admin, err := uc.repo.GetAdminData(ctx, request.Username)
	if err != nil {
		return "", err
	}

	if admin.Password != request.Password {
		return "", errors.New("xato parol kiritdingiz")
	}

	expDuration := time.Duration(uc.cfg.Casbin.AccessTokenTimeOut) * time.Second
	expTime := time.Now().Add(expDuration)

	jwtHandler := tokens.JWTHandler{
		Sub:       admin.Id,
		Iss:       time.Now().String(),
		Exp:       expTime.String(),
		Role:      "admin",
		SigninKey: uc.cfg.Casbin.SigningKey,
		Timeout:   uc.cfg.Casbin.AccessTokenTimeOut,
	}

	accessToken, _, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (uc *AuthUseCase) SuperAdminLogin(ctx context.Context, request *entity.SuperAdmin) (string, error) {
	admin, err := uc.repo.GetSuperAdminData(ctx, request.PhoneNumber)
	if err != nil {
		return "", err
	}

	if admin.Password != request.Password {
		return "", errors.New("xato parol kiritdingiz")
	}

	expDuration := time.Duration(uc.cfg.Casbin.AccessTokenTimeOut) * time.Second
	expTime := time.Now().Add(expDuration)

	jwtHandler := tokens.JWTHandler{
		Sub:       admin.Id,
		Iss:       time.Now().String(),
		Exp:       expTime.String(),
		Role:      "super-admin",
		SigninKey: uc.cfg.Casbin.SigningKey,
		Timeout:   uc.cfg.Casbin.AccessTokenTimeOut,
	}

	accessToken, _, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (uc *AuthUseCase) CreateAdmin(ctx context.Context, admin *entity.Admin) error {
	if err := uc.repo.CreateAdmin(ctx, admin); err != nil {
		return err
	}

	return nil
}

func (uc *AuthUseCase) DeleteAdmin(ctx context.Context, id string) error {
	if err := uc.repo.DeleteAdmin(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *AuthUseCase) GetAllAdmins(ctx context.Context) ([]entity.Admin, error) {
	admins, err := uc.repo.GetAllAdmins(ctx)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
