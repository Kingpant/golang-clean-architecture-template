package repository

import (
	"context"

	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/model"
	"github.com/uptrace/bun"
)

type UserPGRepository struct {
	db *bun.DB
}

func NewUserPGRepository(db *bun.DB) *UserPGRepository {
	return &UserPGRepository{db: db}
}

func (r *UserPGRepository) FindAll(ctx context.Context) ([]*model.User, error) {
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

func (r *UserPGRepository) UpdateOneEmailByID(ctx context.Context, id, email string) error {
	model := new(model.User)
	_, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", id).
		Set("email = ?", email).
		Set("updated_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	return err
}
