package repository

type UserRepository struct {
	  
}

func(r *UserRepository) CreateUser() {
	// create user
}

func(r *UserRepository) FindById() {
	// 先找cache
	// 找不到再找daoao
	// 找到后存入cache
}