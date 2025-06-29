package repository

import (
	"context"

	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/model"
)

//go:generate mockgen -source=user_repository.go -destination=mocks/mock_user_repository.go -package=mocks
type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	UpdateOneEmailByID(ctx context.Context, id, email string) error
}
