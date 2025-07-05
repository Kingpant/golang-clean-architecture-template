package repository

import (
	"context"
	"database/sql"

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
	if err := r.db.NewSelect().Model(&users).Relation("UserLogs").Scan(ctx); err != nil {
		return nil, err
	}

	var usersModel []*model.User
	for _, user := range users {
		usersModel = append(usersModel, &model.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return usersModel, nil
}

func (r *UserPGRepository) Create(ctx context.Context, user *model.User) error {
	if terr := r.db.RunInTx(
		ctx,
		&sql.TxOptions{Isolation: sql.LevelSerializable},
		func(ctx context.Context, tx bun.Tx) error {
			dbUser := &dbmodel.User{
				Name:  user.Name,
				Email: user.Email,
			}

			if _, err := tx.NewInsert().Model(dbUser).Exec(ctx); err != nil {
				return err
			}

			user.ID = dbUser.ID
			user.CreatedAt = dbUser.CreatedAt

			userLog := &dbmodel.UserLog{
				UserID: dbUser.ID,
				Action: dbmodel.UserLogActionTypeCreate,
			}

			if _, err := tx.NewInsert().Model(userLog).Exec(ctx); err != nil {
				return err
			}

			return nil
		},
	); terr != nil {
		return terr
	}

	return nil
}

func (r *UserPGRepository) FindThenUpdateOneEmailByID(ctx context.Context, id, email string) error {
	return r.db.RunInTx(
		ctx,
		&sql.TxOptions{Isolation: sql.LevelSerializable},
		func(ctx context.Context, tx bun.Tx) error {
			user := &dbmodel.User{ID: id}
			if err := tx.NewSelect().Model(user).Where("id = ?", id).For("UPDATE").Scan(ctx); err != nil {
				return err
			}

			user.Email = email
			if _, err := tx.NewUpdate().Model(user).WherePK().Set("updated_at = NOW()").Exec(ctx); err != nil {
				return err
			}

			userLog := &dbmodel.UserLog{
				UserID: user.ID,
				Action: dbmodel.UserLogActionTypeUpdate,
			}
			if _, err := tx.NewInsert().Model(userLog).Exec(ctx); err != nil {
				return err
			}

			return nil
		},
	)
}
