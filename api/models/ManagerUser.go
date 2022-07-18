package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// gorm => a database ORM
type ManagerUser struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Phone     string    `gorm:"size:100;not null;unique" json:"phone"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// bcrypt => a hash tool, for hash password and verify password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *ManagerUser) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *ManagerUser) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *ManagerUser) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Phone == "" {
			return errors.New("Required Phone")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		// if u.Phone == "" {
		// 	return errors.New("Required Phone")
		// }
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *ManagerUser) SaveManagerUser(db *gorm.DB) (*ManagerUser, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &ManagerUser{}, err
	}
	return u, nil
}

func (u *ManagerUser) FindAllManagerUsers(db *gorm.DB) (*[]ManagerUser, error) {
	var err error
	users := []ManagerUser{}
	err = db.Debug().Model(&ManagerUser{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]ManagerUser{}, err
	}
	return &users, err
}

func (u *ManagerUser) FindManagerUserByID(db *gorm.DB, uid uint32) (*ManagerUser, error) {
	var err error
	err = db.Debug().Model(ManagerUser{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &ManagerUser{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &ManagerUser{}, errors.New("ManagerUser Not Found")
	}
	return u, err
}

func (u *ManagerUser) UpdateAManagerUser(db *gorm.DB, uid uint32) (*ManagerUser, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&ManagerUser{}).Where("id = ?", uid).Take(&ManagerUser{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"Name":      u.Name,
			"email":     u.Email,
			"phone":     u.Phone,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &ManagerUser{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&ManagerUser{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &ManagerUser{}, err
	}
	return u, nil
}

func (u *ManagerUser) DeleteAManagerUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&ManagerUser{}).Where("id = ?", uid).Take(&ManagerUser{}).Delete(&ManagerUser{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
