package data

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64  `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Img       string `gorm:"column:img" json:"img"`
}

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db}
}

func (d *UserData) Read(id int64) (*User, error) {
	user := User{}
	if err := d.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserData) ReadAll() ([]*User, error) {
	users := make([]*User, 0)
	if err := d.db.Find(&users).Error; err != nil {
		return []*User{}, err
	}
	return users, nil
}

func (d *UserData) UserPaginator(page, limit int64) (*pagination.Paginator, error) {
	var users []*User
	list := d.db.Find(&users)
	if list.Error != nil {
		return nil, list.Error
	}
	paginator := pagination.Paging(&pagination.Param{
		DB:      list,
		Page:    int(page),
		Limit:   int(limit),
		OrderBy: []string{"id asc"},
	}, &users)
	return paginator, nil
}

func (d *UserData) Create(user *User) (int64, error) {
	if err := d.db.Create(user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}

func (d *UserData) Update(new *User) (*User, error) {
	if err := d.db.Save(&new).Error; err != nil {
		return nil, err
	}
	return new, nil
}

func (d *UserData) Delete(id int64) (int64, error) {
	if _, err := d.Read(id); err != nil {
		return -1, err
	}
	if err := d.db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return -1, err
	}
	return id, nil
}
