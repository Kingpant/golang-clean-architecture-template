package repository

import (
	"context"

	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/model"
	dbmodel "github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/db/model"
	"github.com/uptrace/bun"
)

type UserPGRepository struct {
	db *bun.DB
}

func NewUserPGRepository(db *bun.DB) *UserPGRepository {
	return &UserPGRepository{db: db}
}

func (r *UserPGRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*dbmodel.User
	if err := r.db.NewSelect().Model(&users).Scan(ctx); err != nil {
		return nil, err
	}

	var usersModel []*model.User
	for _, user := range users {
		usersModel = append(usersModel, &model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return usersModel, nil
}

func (r *UserPGRepository) Create(ctx context.Context, name, email string) (string, error) {
	userModel := &dbmodel.User{
		Name:  name,
		Email: email,
	}

	_, err := r.db.NewInsert().Model(userModel).Returning("id").Exec(ctx)
	return userModel.ID, err
}

func (r *UserPGRepository) UpdateOneEmailByID(ctx context.Context, id, email string) error {
	model := new(dbmodel.User)
	_, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", id).
		Set("email = ?", email).
		Set("updated_at = NOW()").
		Where("id = ?", id).
		Exec(ctx)
	return err
}
