package usecase

import (
	"context"
	"errors"

	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/model"
	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/repository"
	"go.uber.org/zap"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]string, []string, error)
	CreateUser(ctx context.Context, name, email string) (string, error)
	UpdateUserEmail(ctx context.Context, id string, email string) error
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

func (u *userUsecase) GetUsers(ctx context.Context) ([]string, []string, error) {
	userModels, err := u.userRepo.FindAll(ctx)
	if err != nil {
		u.logger.Errorw("failed to get users", "error", err)
		return nil, nil, err
	}

	users := []string{}
	userIDs := []string{}
	for _, user := range userModels {
		users = append(users, user.Name)
		userIDs = append(userIDs, user.ID)
	}

	return users, userIDs, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, name, email string) (string, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	if !user.IsValidEmail() || !user.IsValidName() {
		u.logger.Errorw("invalid user data", "name", name, "email", email)
		return "", errors.New("invalid user data")
	}

	id, err := u.userRepo.Create(ctx, name, email)
	if err != nil {
		u.logger.Errorw("failed to create user", "error", err)
		return "", err
	}

	return id, nil
}

func (u *userUsecase) UpdateUserEmail(ctx context.Context, id string, email string) error {
	err := u.userRepo.FindThenUpdateOneEmailByID(ctx, id, email)
	if err != nil {
		u.logger.Errorw("failed to update user", "error", err)
		return err
	}

	return nil
}
