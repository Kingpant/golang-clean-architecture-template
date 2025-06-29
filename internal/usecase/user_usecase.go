package usecase

import (
	"context"

	"github.com/Kingpant/tipster/internal/domain/model"
	"github.com/Kingpant/tipster/internal/domain/repository"
	"go.uber.org/zap"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]string, error)
	CreateUser(ctx context.Context, name, email string) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository

	logger *zap.SugaredLogger
}

func NewUserUseCase(userRepo repository.UserRepository, logger *zap.SugaredLogger) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]string, error) {
	userModels, err := u.userRepo.GetAll(ctx)
	if err != nil {
		u.logger.Errorw("failed to get users", "error", err)
		return nil, err
	}

	users := []string{}
	for _, user := range userModels {
		users = append(users, user.Name)
	}

	return users, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, name, email string) (string, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}

	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		u.logger.Errorw("failed to create user", "error", err)
		return "", err
	}

	return id, nil
}
