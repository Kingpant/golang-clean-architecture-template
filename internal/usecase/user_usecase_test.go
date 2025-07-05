package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/model"
	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/repository"
	"github.com/Kingpant/golang-clean-architecture-template/internal/domain/repository/mocks"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func Test_userUsecase_CreateUser(t *testing.T) {
	type fields struct {
		userRepo repository.UserRepository
		logger   *zap.SugaredLogger
	}
	type args struct {
		ctx   context.Context
		name  string
		email string
	}

	// Mock user repository
	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockUserRepo.EXPECT().
		Create(gomock.Any(), "Alice", "alice@example.com").
		Return("12345", nil).
		Times(1)

	mockUserRepo.EXPECT().
		Create(gomock.Any(), "Bob", "bob@example.com").
		Return("", errors.New("user already exists")).
		Times(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CreateUser_Success",
			fields: fields{
				userRepo: mockUserRepo,
				logger:   zap.NewNop().Sugar(),
			},
			args: args{
				ctx:   context.Background(),
				name:  "Alice",
				email: "alice@example.com",
			},
			want:    "12345",
			wantErr: false,
		},
		{
			name: "CreateUser_Failure",
			fields: fields{
				userRepo: mockUserRepo,
				logger:   zap.NewNop().Sugar(),
			},
			args: args{
				ctx:   context.Background(),
				name:  "Bob",
				email: "bob@example.com",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepo: tt.fields.userRepo,
				logger:   tt.fields.logger,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateUserEmail(t *testing.T) {
	type fields struct {
		userRepo repository.UserRepository
		logger   *zap.SugaredLogger
	}
	type args struct {
		ctx   context.Context
		id    string
		email string
	}

	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockUserRepo.EXPECT().
		FindThenUpdateOneEmailByID(gomock.Any(), "12345", "new_email@example.com").
		Return(nil).
		Times(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "UpdateUserEmail_Success",
			fields: fields{
				userRepo: mockUserRepo,
				logger:   zap.NewNop().Sugar(),
			},
			args: args{
				ctx:   context.Background(),
				id:    "12345",
				email: "new_email@example.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepo: tt.fields.userRepo,
				logger:   tt.fields.logger,
			}
			if err := u.UpdateUserEmail(tt.args.ctx, tt.args.id, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.UpdateUserEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUsecase_GetUsers(t *testing.T) {
	type fields struct {
		userRepo repository.UserRepository
		logger   *zap.SugaredLogger
	}
	type args struct {
		ctx context.Context
	}

	// Mock user repository
	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockUserRepo.EXPECT().
		FindAll(gomock.Any()).
		Return([]*model.User{
			{ID: "1", Name: "Alice", Email: "alice@example.com"},
			{ID: "2", Name: "Bob", Email: "bob@example.com"},
		}, nil).
		Times(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		want1   []string
		wantErr bool
	}{
		{
			name: "GetUsers_Success",
			fields: fields{
				userRepo: mockUserRepo,
				logger:   zap.NewNop().Sugar(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    []string{"Alice", "Bob"},
			want1:   []string{"1", "2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepo: tt.fields.userRepo,
				logger:   tt.fields.logger,
			}
			got, got1, err := u.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUsers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userUsecase.GetUsers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
