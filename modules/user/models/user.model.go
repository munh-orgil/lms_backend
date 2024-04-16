package user_models

import (
	"lms_backend/database"

	"lms_backend/helpers"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	Username       string         `json:"username" gorm:"type:varchar(255)" validate:"required"`
	Password       string         `json:"-"`
	Email          string         `json:"email" gorm:"type:varchar(255)" validate:"required"`
	ProfilePicture string         `json:"profile_picture" gorm:"type:varchar(255)"`
	UserType       string         `json:"user_type" gorm:"type:varchar(50)" validate:"required"`
	Lastname       string         `json:"lastname" gorm:"type:varchar(255)" validate:"required"`
	Firstname      string         `json:"firstname" gorm:"type:varchar(255)" validate:"required"`
	Gender         uint           `json:"gender" validate:"required"`
	BirthDate      data.LocalDate `json:"birth_date" gorm:"type:date" validate:"required"`
	CreatedAt      data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

type ReqUserCreate struct {
	Username       string         `json:"username" validate:"required"`
	Password       string         `json:"password" validate:"required"`
	Email          string         `json:"email" validate:"required"`
	ProfilePicture string         `json:"profile_picture"`
	UserType       string         `json:"user_type" validate:"required"`
	Lastname       string         `json:"lastname" validate:"required"`
	Firstname      string         `json:"firstname" validate:"required"`
	Gender         uint           `json:"gender" validate:"required"`
	BirthDate      data.LocalDate `json:"birth_date" validate:"required"`
}

type ReqUserUpdate struct {
	Id             uint           `json:"id" validate:"required"`
	Username       string         `json:"username"`
	ProfilePicture string         `json:"profile_picture"`
	UserType       string         `json:"user_type"`
	Lastname       string         `json:"lastname"`
	Firstname      string         `json:"firstname"`
	Gender         uint           `json:"gender"`
	BirthDate      data.LocalDate `json:"birth_date"`
}

func UserList(c *fiber.Ctx) (res *data.Pagination[User], err error) {
	var totalRows int64
	users := make([]User, 0)

	db := database.DBconn
	tx := db.Model(User{})
	tx.Count(&totalRows)

	p := data.Paginate[User](c, totalRows)

	if err = tx.Offset(p.Offset).Limit(p.PageSize).Find(&users).Error; err != nil {
		return
	}
	p.Items = users

	res = p

	return
}

func (u *User) Create() error {
	db := database.DBconn
	u.Password = helpers.GeneratePassword(u.Password)
	return db.Create(&u).Error
}

func (u *User) Update() error {
	db := database.DBconn
	return db.Updates(&u).Error
}

func FindUserBy(column string, val interface{}) (u User, err error) {
	db := database.DBconn
	err = db.Where(column+" = ?", val).First(&u).Error
	return
}
