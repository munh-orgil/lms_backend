package subject_models

import (
	"lms_backend/database"
	user_models "lms_backend/modules/user/models"

	"github.com/craftzbay/go_grc/v2/data"
)

type Subject struct {
	Id         uint             `json:"id" gorm:"primaryKey"`
	Name       string           `json:"name"`
	TeacherId  uint             `json:"-"`
	Teacher    user_models.User `json:"teacher" gorm:"foreignKey:TeacherId"`
	Banner     string           `json:"banner"`
	TotalScore float64          `json:"total_score"`
	CreatedAt  data.LocalTime   `json:"created_at" gorm:"autoCreateTime"`
}

func SubjectList(studentId uint) (res []Subject, err error) {
	db := database.DBconn
	ids := make([]uint, 0)
	if err = db.Model(StudentSubject{}).Where("student_id = ?", studentId).Select("id").Scan(&ids).Error; err != nil {
		return
	}
	err = db.Where("id IN ?", ids).Order("name ASC").Preload("Teacher").Find(&res).Error
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
