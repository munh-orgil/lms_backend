package subject_models

import (
	"encoding/json"
	"lms_backend/database"
	"lms_backend/modules/attachment"
	user_models "lms_backend/modules/user/models"
	"lms_backend/utils"
	"sort"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Subject struct {
	Id           uint             `json:"id" gorm:"primaryKey"`
	Name         string           `json:"name"`
	TeacherId    uint             `json:"-"`
	Teacher      user_models.User `json:"teacher" gorm:"foreignKey:TeacherId"`
	Banner       string           `json:"banner"`
	StudentScore float64          `json:"student_score"`
	TotalScore   float64          `json:"total_score"`
	Lectures     []Lecture        `json:"lectures" gorm:"foreignKey:SubjectId"`
	CreatedAt    data.LocalTime   `json:"created_at" gorm:"autoCreateTime"`
}

func (*Subject) BeforeFind(tx *gorm.DB) error {
	return tx.Preload("Teacher").Error
}

func (s *Subject) AfterFind(tx *gorm.DB) error {
	sort.Slice(s.Lectures, func(i2, j int) bool {
		return time.Time(s.Lectures[i2].CreatedAt).Before(time.Time(s.Lectures[j].CreatedAt))
	})
	return nil
}

func SubjectList(studentId uint) (res []Subject, err error) {
	db := database.DBconn
	err = db.Table(utils.GetTableName("subject", "s")).
		Joins("INNER JOIN "+utils.GetTableName("student_subject", "ss")+" ON ss.subject_id = s.id").
		Select("s.*, ss.score as student_score").
		Where("ss.student_id = ?", studentId).Order("name ASC").Preload("Teacher").Preload("Lectures").Find(&res).Error
	return
}

func (s *Subject) Create() error {
	db := database.DBconn
	return db.Create(&s).Error
}

func (s *Subject) Update() error {
	db := database.DBconn
	return db.Updates(&s).Error
}

func (s *Subject) Delete() error {
	db := database.DBconn
	return db.Delete(&s).Error
}

type Lecture struct {
	Id            uint                    `json:"id" gorm:"primaryKey"`
	Title         string                  `json:"title" gorm:"type:varchar(50)"`
	Description   string                  `json:"description"`
	AttachmentIds datatypes.JSON          `json:"attachment_ids"`
	Attachments   []attachment.Attachment `json:"attachments" gorm:"-"`
	SubjectId     uint                    `json:"subject_id"`
	CreatedAt     data.LocalTime          `json:"created_at" gorm:"autoCreateDate"`
}

func (l *Lecture) AfterFind(tx *gorm.DB) error {
	attachmentIds := make([]uint, 0)
	json.Unmarshal(l.AttachmentIds, &attachmentIds)
	return tx.Where("id IN ?", attachmentIds).Find(&l.Attachments).Error
}
