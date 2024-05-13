package schedule_models

import (
	"lms_backend/database"
	subject_models "lms_backend/modules/subject/models"
	"time"
)

type Schedule struct {
	Id        uint                   `json:"id" gorm:"primaryKey"`
	DayOfWeek uint                   `json:"day_of_week"`
	SubjectId uint                   `json:"subject_id"`
	Subject   subject_models.Subject `json:"subject" gorm:"foreignKey:SubjectId"`
	StartAt   time.Time              `json:"start_at"`
	EndAt     time.Time              `json:"end_at"`
}

type ScheduleByGroup struct {
	DayOfWeek uint       `json:"day_of_week"`
	Schedules []Schedule `json:"schedules"`
}

func ScheduleList() (res []ScheduleByGroup, err error) {
	db := database.DBconn

	schedules := []Schedule{}
	if err = db.Order("start_at ASC").Preload("Subject").Preload("Subject.Teacher").Find(&schedules).Error; err != nil {
		return
	}

	res = make([]ScheduleByGroup, 7)
	for i := range 7 {
		res[i].DayOfWeek = uint(i)
	}
	for _, s := range schedules {
		res[s.DayOfWeek].Schedules = append(res[s.DayOfWeek].Schedules, s)
	}

	return
}
