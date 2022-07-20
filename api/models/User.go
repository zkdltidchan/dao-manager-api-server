package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zkdltidchan/dao-manager-api-server/api/responses"
)

// gorm => a database ORM
type User struct {
	UserID          string    `gorm:"size:20;primary_key;notnull" json:"user_id"`
	UserName        string    `gorm:"size:10;not null;unique" json:"user_name"`
	UserNick        string    `gorm:"size:30;not null;unique" json:"user_nick"`
	UserPw          string    `gorm:"size:20;not null;" json:"user_password"`
	UserKBank       string    `gorm:"size:20;not null;unique" json:"user_kbank"`
	UserPhone       string    `gorm:"size:15;not null;unique" json:"user_phone"`
	UserEmail       string    `gorm:"size:40;not null;unique" json:"user_email"`
	UserJoinDate    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_join_date"`
	UserUpdataTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_update_time"`
	UserRecentLogin time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_recent_login"`
}

// TODO, user data response
// type UserDat struct {
// 	UserID          string    `gorm:"size:20;primary_key;notnull" json:"user_id"`
// 	UserName        string    `gorm:"size:10;not null;unique" json:"user_name"`
// 	UserNick        string    `gorm:"size:30;not null;unique" json:"user_nick"`
// 	UserKBank       string    `gorm:"size:20;not null;unique" json:"user_kbank"`
// 	UserPhone       string    `gorm:"size:15;not null;unique" json:"user_phone"`
// 	UserEmail       string    `gorm:"size:40;not null;unique" json:"user_email"`
// 	UserJoinDate    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_join_date"`
// 	UserUpdataTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_update_time"`
// 	UserRecentLogin time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"user_recent_login"`
// }

type UserListResponse struct {
	responses.FeatchListResponse
	Data []User `json:"data"`
}

type UserParameter struct {
	PageIndex int `schema:"page_index"`
	Size      int `schema:"size"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.UserPw)
	if err != nil {
		return err
	}
	u.UserPw = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.UserID = "0"
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.UserEmail = html.EscapeString(strings.TrimSpace(u.UserEmail))
	u.UserJoinDate = time.Now()
	u.UserUpdataTime = time.Now()
}

func (u *User) FindAllUsers(db *gorm.DB, userParameter UserParameter) (*UserListResponse, error) {
	// func (u *User) FindAllUsers(db *gorm.DB, userParameter UserParameter) (*[]User, error) {
	var err error
	responses := UserListResponse{}
	// get all users count with filter,TODO : filter
	err = db.Debug().Model(&User{}).Count(&responses.Total).Error
	if err != nil {
		return &responses, err
	}
	responses.PageCounts = GetPages(responses.Total, userParameter.Size)
	limit := GetLimit(userParameter.Size)
	offSet := GetOffSet(userParameter.PageIndex, limit, responses.PageCounts)

	// get users
	err = db.Debug().Model(&User{}).Limit(limit).Offset(offSet).Find(&responses.Data).Error
	if err != nil {
		return &responses, err
	}
	responses.Size = len(responses.Data)
	responses.PageIndex = GetCurrentPage(offSet)
	return &responses, err

	// members := []User{}
	// err = db.Debug().Model(&User{}).Limit(userParameter.Size).Find(&members).Error
	// err = db.Debug().Model(&User{}).Limit(100).Find(&members).Error
	// if err != nil {
	// 	return &[]User{}, err
	// }
	// responses.Data = members
	// return &members, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
