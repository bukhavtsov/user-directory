package data

import (
	"github.com/jinzhu/gorm"
)

type UserData interface {
	Create(user *User) (int64, error)
	Read(id int64) (*User, error)
	ReadAll() ([]*User, error)
	Update(user *User) (*User, error)
	Delete(id int64) error
}

type User struct {
	Id        int64  `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	ImgName   string `gorm:"column:img_name"`
	Img       []byte `gorm:"column:img"`
}

type userData struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userData {
	return &userData{db}
}

func (d *userData) Read(id int64) (*User, error) {
	user := User{}
	if err := d.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *userData) ReadAll() ([]*User, error) {
	users := make([]*User, 0)
	if err := d.db.Find(&users).Error; err != nil {
		return []*User{}, err
	}
	return users, nil
}

func (d *userData) Create(user *User) (int64, error) {
	if err := d.db.Create(user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}

func (d *userData) Update(user *User) (*User, error) {
	if err := d.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d *userData) Delete(id int64) error {
	if err := d.db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
