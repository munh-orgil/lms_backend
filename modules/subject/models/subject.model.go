package subject_models

import (
	"lms_backend/database"
	user_models "lms_backend/modules/user/models"
	"sort"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
)

type Subject struct {
	Id         uint             `json:"id" gorm:"primaryKey"`
	Name       string           `json:"name"`
	TeacherId  uint             `json:"-"`
	Teacher    user_models.User `json:"teacher" gorm:"foreignKey:TeacherId"`
	Banner     string           `json:"banner"`
	TotalScore float64          `json:"total_score"`
	Lectures   []Lecture        `json:"lectures" gorm:"foreignKey:SubjectId"`
	CreatedAt  data.LocalTime   `json:"created_at" gorm:"autoCreateTime"`
}

func SubjectList(studentId uint) (res []Subject, err error) {
	db := database.DBconn
	ids := make([]uint, 0)
	if err = db.Model(StudentSubject{}).Where("student_id = ?", studentId).Select("id").Scan(&ids).Error; err != nil {
		return
	}
	err = db.Where("id IN ?", ids).Order("name ASC").Preload("Teacher").Preload("Lectures").Find(&res).Error
	for i := range res {
		sort.Slice(res[i].Lectures, func(i2, j int) bool {
			return time.Time(res[i].Lectures[i2].CreatedAt).After(time.Time(res[i].Lectures[j].CreatedAt))
		})
	}
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
	Id          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(50)"`
	Attachments datatypes.JSON `json:"attachments"`
	SubjectId   uint           `json:"subject_id"`
	CreatedAt   data.LocalTime `json:"created_at" gorm:"autoCreateDate"`
}
