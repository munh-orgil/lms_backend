package subject_models

import (
	"lms_backend/database"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/gorm"
)

type Exam struct {
	Id           uint           `json:"id" gorm:"primaryKey"`
	SubjectId    uint           `json:"subject_id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Source       string         `json:"source"`
	ActiveStatus uint           `json:"active_status" gorm:"-"` // 0:not started; 1:active; 2:over
	StartAt      data.LocalTime `json:"start_at"`
	EndAt        data.LocalTime `json:"end_at"`
	Duration     uint           `json:"duration" gorm:"-"` // in minutes
	StudentScore uint           `json:"student_score"`
	TotalScore   uint           `json:"total_score"`
	CreatedAt    data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

func RunMigrations() {
	database.DBconn.AutoMigrate(Exam{})
}

func (e *Exam) AfterFind(tx *gorm.DB) error {
	e.Duration = uint(time.Time(e.EndAt).Sub(time.Time(e.StartAt)).Minutes())
	switch {
	case time.Time(e.StartAt).After(time.Now()):
		e.ActiveStatus = 0
	case time.Time(e.StartAt).Before(time.Now()) && time.Time(e.EndAt).After(time.Now()):
		e.ActiveStatus = 1
	case time.Time(e.EndAt).Before(time.Now()):
		e.ActiveStatus = 2
	}
	return nil
}

func ExamList(subjectId uint) (res []Exam, err error) {
	db := database.DBconn
	if err = db.Model(Exam{}).Where("subject_id = ?", subjectId).Order("start_at ASC").Find(&res).Error; err != nil {
		return
	}
	return
}

func (e *Exam) Create() error {
	db := database.DBconn
	return db.Create(&e).Error
}

func (e *Exam) Update() error {
	db := database.DBconn
	return db.Updates(&e).Error
}

func (e *Exam) Delete() error {
	db := database.DBconn
	return db.Delete(&e).Error
}
