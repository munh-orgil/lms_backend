package subject_models

import (
	"lms_backend/database"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
)

type StudentSubject struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	StudentId uint    `json:"-"`
	SubjectId uint    `json:"-"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectId"`
	Score     float64 `json:"score"`
	Tasks     Tasks   `json:"tasks" gorm:"-"`
}

type Tasks struct {
	Late []Task        `json:"late" gorm:"-"`
	Due  []Task        `json:"due" gorm:"-"`
	Done []StudentTask `json:"done" gorm:"-"`
}

type StudentTask struct {
	Id         uint           `json:"id" gorm:"primaryKey"`
	Attachment string         `json:"attachment"`
	StudentId  uint           `json:"student_id"`
	TaskId     uint           `json:"task_id"`
	Task       Task           `json:"task" gorm:"foreignKey:TaskId"`
	Score      float64        `json:"score"`
	CreatedAt  data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

func GetSubject(studentId, subjectId uint) (res *StudentSubject, err error) {
	db := database.DBconn

	if err = db.Where("student_id = ?", studentId).Where("subject_id = ?", subjectId).
		Preload("Subject").Preload("Subject.Lectures").First(&res).Error; err != nil {
		return
	}

	subjectTaskIds := make([]uint, 0)
	if err = db.Model(Task{}).Where("subject_id = ?", subjectId).Select("id").Scan(&subjectTaskIds).Error; err != nil {
		return
	}

	studentTasks := make([]StudentTask, 0)
	if err = db.Where("student_id = ?", studentId).Where("task_id IN ?", subjectTaskIds).Preload("Task").Find(&studentTasks).Error; err != nil {
		return
	}
	doneTaskIds := make([]uint, 0)
	for _, st := range studentTasks {
		doneTaskIds = append(doneTaskIds, st.Id)
	}

	tasks := make([]Task, 0)
	if err = db.Where("subject_id = ?", subjectId).Where("id NOT IN ?", doneTaskIds).Find(&tasks).Error; err != nil {
		return
	}

	res.Tasks.Done = studentTasks
	for _, t := range tasks {
		if time.Time(t.Due).After(time.Now()) {
			res.Tasks.Late = append(res.Tasks.Late, t)
		} else {
			res.Tasks.Due = append(res.Tasks.Due, t)
		}
	}

	return
}
