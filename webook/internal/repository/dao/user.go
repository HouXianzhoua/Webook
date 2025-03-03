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

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	// find user by email
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
} 

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