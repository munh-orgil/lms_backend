package schedule_models

import (
	"lms_backend/database"
	subject_models "lms_backend/modules/subject/models"
	"strings"
	"time"

	"github.com/craftzbay/go_grc/v2/converter"
	"gorm.io/gorm"
)

type Schedule struct {
	Id        uint                   `json:"id" gorm:"primaryKey"`
	DayOfWeek uint                   `json:"day_of_week"`
	SubjectId uint                   `json:"subject_id"`
	Subject   subject_models.Subject `json:"subject" gorm:"foreignKey:SubjectId"`
	StartAt   string                 `json:"start_at"`
	EndAt     string                 `json:"end_at"`
	Duration  uint                   `json:"duration" gorm:"-"`
	IsActive  bool                   `json:"is_active" gorm:"-"`
}

func (s *Schedule) AfterFind(tx *gorm.DB) error {
	start := strings.Split(s.StartAt, ":")
	end := strings.Split(s.EndAt, ":")
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), converter.StringToInt(start[0]), converter.StringToInt(start[1]), 0, 0, time.Local)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), converter.StringToInt(end[0]), converter.StringToInt(end[1]), 0, 0, time.Local)
	s.Duration = uint(endTime.Sub(startTime).Minutes())
	if startTime.Before(time.Now()) && endTime.After(time.Now()) {
		s.IsActive = true
	}
	return nil
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
		res[i].Schedules = make([]Schedule, 0)
	}
	for _, s := range schedules {
		res[s.DayOfWeek].Schedules = append(res[s.DayOfWeek].Schedules, s)
	}

	return
}
