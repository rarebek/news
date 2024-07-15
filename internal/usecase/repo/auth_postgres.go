package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tarkib.uz/internal/entity"
	"tarkib.uz/pkg/postgres"
)

type AuthRepo struct {
	*postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg}
}

func (a *AuthRepo) GetAdminData(ctx context.Context, PhoneNumber string) (*entity.Admin, error) {
	var adminPostgres entity.Admin
	sql, args, err := a.Builder.Select("id, phone_number, password").
		From("admins").
		Where(squirrel.Eq{
			"phone_number": PhoneNumber,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	row := a.Pool.QueryRow(ctx, sql, args...)

	if err = row.Scan(&adminPostgres.Id, &adminPostgres.PhoneNumber, &adminPostgres.Password); err != nil {
		return nil, err
	}

	return &adminPostgres, nil
}

func (a *AuthRepo) GetSuperAdminData(ctx context.Context, PhoneNumber string) (*entity.Admin, error) {
	var adminPostgres entity.Admin
	sql, args, err := a.Builder.Select("id, phone_number, password").
		From("superadmins").
		Where(squirrel.Eq{
			"phone_number": PhoneNumber,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	row := a.Pool.QueryRow(ctx, sql, args...)

	if err = row.Scan(&adminPostgres.Id, &adminPostgres.PhoneNumber, &adminPostgres.Password); err != nil {
		return nil, err
	}

	return &adminPostgres, nil
}

func (a *AuthRepo) CreateAdmin(ctx context.Context, admin *entity.Admin) error {
	data := map[string]interface{}{
		"id":           uuid.NewString(),
		"phone_number": admin.PhoneNumber,
		"password":     admin.Password,
	}

	sql, args, err := a.Builder.Insert("admins").
		SetMap(data).ToSql()
	if err != nil {
		return err
	}

	if _, err = a.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
