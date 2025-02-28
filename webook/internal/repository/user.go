package repository

import (
	"context"

	"example.com/webook/internal/domain"
	"example.com/webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
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