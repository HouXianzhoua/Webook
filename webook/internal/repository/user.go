package repository

import (
	"context"
	// "errors"

	"example.com/webook/internal/domain"
	"example.com/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound=dao.ErrUserNotFound 
)
type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// find user by email
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err 
	}
		
	return domain.User{ 
		ID: u.ID,
		Email: u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	// create user
	return r.dao.Insert(ctx,dao.User{

		Email:    u.Email, 
		Password: u.Password,
	})
}

func (r *UserRepository) FindById() {
	// 先找cache
	// 找不到再找daoao
	// 找到后存入cache
}