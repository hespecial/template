package repo

import (
	"context"
	"template/internal/model"
	"template/internal/repo/dao"
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id uint) (*model.User, error)
}

type userRepo struct {
	dao dao.UserDao
}

func NewUserRepo(dao dao.UserDao) UserRepo {
	return &userRepo{
		dao: dao,
	}
}

func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return u.dao.GetUserByUsername(ctx, username)
}

func (u *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	return u.dao.CreateUser(ctx, user)
}

func (u *userRepo) GetUserById(ctx context.Context, id uint) (*model.User, error) {
	return u.dao.GetUserById(ctx, id)
}
