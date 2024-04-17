package subject_models

import (
	"fmt"
	"lms_backend/database"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	SubjectId   uint           `json:"-"`
	Subject     Subject        `json:"subject" gorm:"SubjectId"`
	Score       float64        `json:"score"`
	Due         data.LocalTime `json:"due"`
}

func TaskList(c *fiber.Ctx) (res []Task, err error) {
	db := database.DBconn
	tx := db.Model(Task{})
	err = tx.Order("due DESC").Preload("Subject").Find(&res).Error
	fmt.Printf("res: %v\n", res)
	return
}

func (t *Task) Create() error {
	db := database.DBconn
	return db.Create(&t).Error
}

func (t *Task) Update() error {
	db := database.DBconn
	return db.Updates(&t).Error
}

func (t *Task) Delete() error {
	db := database.DBconn
	return db.Delete(&t).Error
}
