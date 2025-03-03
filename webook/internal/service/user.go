package service

import (
	"context"
	"errors"

	"example.com/webook/internal/domain"
	"example.com/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("invalid user or password")

type UserService struct{
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService{	
	return &UserService{
		repo:repo,
	}
}

func(svc *UserService) Login (ctx context.Context,email,password string)(domain.User,error){
	u,err:=svc.repo.FindByEmail(ctx,email)
	if err==repository.ErrUserNotFound{
		return domain.User{},ErrInvalidUserOrPassword
	} 
	if err!=nil{
		return domain.User{},err	
	}  
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
    	return domain.User{},ErrInvalidUserOrPassword
	}
	return u,nil	
} 

func(svc *UserService) Signup(ctx context.Context,u domain.User)error{
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {	
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx,u)
} 