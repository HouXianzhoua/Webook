package domain

import "time"

//User领域对象，是DDD的聚合根/entity
//BO（Business Object）业务对象
type User struct {//业务意义
	ID       int64
	Username string		
	Password string
	Email    string
	Ctime    time.Time
}