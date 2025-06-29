package repository

import (
	"context"

	"github.com/Kingpant/golang-template/internal/domain/model"
	"github.com/uptrace/bun"
)

type UserPGRepository struct {
	db *bun.DB
}

func NewUserPGRepository(db *bun.DB) *UserPGRepository {
	return &UserPGRepository{db: db}
}

func (r *UserPGRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPGRepository) Create(ctx context.Context, user *model.User) (string, error) {
	_, err := r.db.NewInsert().Model(user).Returning("id").Exec(ctx)
	return user.ID, err
}
