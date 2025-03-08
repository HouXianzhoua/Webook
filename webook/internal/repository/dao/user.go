package dao

import (
	"context"
	"errors"
	"time"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail=errors.New("duplicate email")
	ErrUserNotFound=gorm.ErrRecordNotFound
)
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// FindByEmail retrieves a user from the database by their email address.
// It returns the User entity if found, or an error if the user does not exist
// or if there is a problem with the database query.

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	// find user by email
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
} 

// Insert inserts a user into the database.
// It returns ErrUserDuplicateEmail if the email address is already in use,
// or an error if there is a problem with the database query.
func (dao *UserDao) Insert(ctx context.Context, u User) error {
	// insert user
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err:=dao.db.WithContext(ctx).Create(&u).Error 
	if mysql,ok:=err.(*mysql.MySQLError);ok{
		const uniqueConflictsErrorNo uint16 = 1062 
		if mysql.Number==uniqueConflictsErrorNo{
			return ErrUserDuplicateEmail
		}
	}	
	return err
	// return dao.db.WithContext(ctx).Create(&u).Error
}

// User represents a user entity User直接对应数据库表结构
// 有些人叫做model，有些人叫做entity（数据库层面），有些人叫做PO（Persistence Object）
type User struct {//数据库意义//数据库存储
	ID       int64 `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Ctime int64
	Utime int64
}