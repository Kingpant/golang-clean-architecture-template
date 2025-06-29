package repository

import (
	"context"

	"github.com/Kingpant/golang-template/internal/domain/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
}
