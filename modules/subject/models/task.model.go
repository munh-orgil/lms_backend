package subject_models

import (
	"encoding/json"
	"lms_backend/database"
	"lms_backend/modules/attachment"
	"lms_backend/session"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	Id            uint                    `json:"id" gorm:"primaryKey"`
	Title         string                  `json:"title"`
	Description   string                  `json:"description"`
	SubjectId     uint                    `json:"-"`
	Subject       Subject                 `json:"subject" gorm:"SubjectId"`
	Score         float64                 `json:"score"`
	Due           data.LocalTime          `json:"due"`
	Type          string                  `json:"type"`
	AttachmentIds datatypes.JSON          `json:"attachment_ids"`
	Attachments   []attachment.Attachment `json:"attachments" gorm:"-"`
	CreatedAt     data.LocalTime          `json:"created_at" gorm:"autoCreateTime"`
}

func (t *Task) AfterFind(tx *gorm.DB) error {
	attachmentIds := make([]uint, 0)
	json.Unmarshal(t.AttachmentIds, &attachmentIds)
	return tx.Where("id IN ?", attachmentIds).Find(&t.Attachments).Error
}

func TaskList(c *fiber.Ctx) (res []Task, err error) {
	db := database.DBconn
	subjectId := c.QueryInt("subject_id")
	tx := db.Model(Task{})
	studentId := session.GetTokenInfo(c).GetUserId()
	doneTaskIds := []uint{}

	if err = db.Model(StudentTask{}).Distinct("task_id").Where("student_id = ?", studentId).Scan(&doneTaskIds).Error; err != nil {
		return
	}
	if len(doneTaskIds) > 0 {
		tx.Where("id not in ?", doneTaskIds)
	}
	if subjectId > 0 {
		tx.Where("subject_id = ?", subjectId)
	}
	err = tx.Order("due DESC").Preload("Subject").Preload("Subject.Teacher").Find(&res).Error
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
