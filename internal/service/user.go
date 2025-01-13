package service

import (
	"context"
	"template/internal/model"
	"template/internal/repo"
	"template/util"
	"time"
)

type UserInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, username, password string) error

	GetUserInfo(ctx context.Context, id uint) (*UserInfo, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}

func (s *userService) CreateUser(ctx context.Context, username, password string) error {
	user := &model.User{
		Username:  username,
		Password:  util.Encrypt(password),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUserInfo(ctx context.Context, id uint) (*UserInfo, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		Id:        id,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
