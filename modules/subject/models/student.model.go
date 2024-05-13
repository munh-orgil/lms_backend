package subject_models

import (
	"encoding/json"
	"lms_backend/database"
	"lms_backend/modules/attachment"
	"sort"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StudentSubject struct {
	Id          uint    `json:"id" gorm:"primaryKey"`
	StudentId   uint    `json:"-"`
	SubjectId   uint    `json:"-"`
	Subject     Subject `json:"subject" gorm:"foreignKey:SubjectId"`
	Score       float64 `json:"score"`
	Tasks       Tasks   `json:"tasks" gorm:"-"`
	Assignments Tasks   `json:"assignments" gorm:"-"`
}

type Tasks struct {
	Late []Task        `json:"late" gorm:"-"`
	Due  []Task        `json:"due" gorm:"-"`
	Done []StudentTask `json:"done" gorm:"-"`
}

type StudentTask struct {
	Id            uint                    `json:"id" gorm:"primaryKey"`
	AttachmentIds datatypes.JSON          `json:"attachement_ids"`
	Attachments   []attachment.Attachment `json:"attachments" gorm:"-"`
	StudentId     uint                    `json:"student_id"`
	TaskId        uint                    `json:"task_id"`
	Task          Task                    `json:"task" gorm:"foreignKey:TaskId"`
	Score         float64                 `json:"score"`
	CreatedAt     data.LocalTime          `json:"created_at" gorm:"autoCreateTime"`
}

func (st *StudentTask) AfterFind(tx *gorm.DB) error {
	ids := []uint{}
	json.Unmarshal(st.AttachmentIds, &ids)
	return tx.Where("id IN ?", ids).Find(&st.Attachments).Error
}

func GetSubject(studentId, subjectId uint) (res *StudentSubject, err error) {
	db := database.DBconn

	if err = db.Where("student_id = ?", studentId).Where("subject_id = ?", subjectId).
		Preload("Subject").Preload("Subject.Teacher").Preload("Subject.Lectures").First(&res).Error; err != nil {
		return
	}

	subjectTaskIds := make([]uint, 0)
	if err = db.Model(Task{}).Where("subject_id = ?", subjectId).Select("id").Scan(&subjectTaskIds).Error; err != nil {
		return
	}

	studentTasks := make([]StudentTask, 0)
	if err = db.Where("student_id = ?", studentId).Where("task_id IN ?", subjectTaskIds).Preload("Task").Preload("Task.Subject").Preload("Task.Subject.Teacher").Find(&studentTasks).Error; err != nil {
		return
	}
	doneTaskIds := make([]uint, 0)
	for _, st := range studentTasks {
		doneTaskIds = append(doneTaskIds, st.TaskId)
	}

	tasks := make([]Task, 0)
	tx := db.Where("subject_id = ?", subjectId)
	if len(doneTaskIds) > 0 {
		tx.Where("id NOT IN ?", doneTaskIds)
	}
	if err = tx.Find(&tasks).Error; err != nil {
		return
	}

	for _, st := range studentTasks {
		if st.Task.Type == "assignment" {
			res.Assignments.Done = append(res.Assignments.Done, st)
		} else {
			res.Tasks.Done = append(res.Tasks.Done, st)
		}
	}
	for _, t := range tasks {
		if t.Type == "assignment" {
			if time.Time(t.Due).After(time.Now()) {
				res.Assignments.Due = append(res.Assignments.Due, t)
			} else {
				res.Assignments.Late = append(res.Assignments.Late, t)
			}
		} else {
			if time.Time(t.Due).After(time.Now()) {
				res.Tasks.Due = append(res.Tasks.Due, t)
			} else {
				res.Tasks.Late = append(res.Tasks.Late, t)
			}
		}
	}
	SortTasks(&res.Tasks)
	return
}

func GetTasks(studentId uint) (res *Tasks, err error) {
	db := database.DBconn
	res = new(Tasks)

	studentTasks := make([]StudentTask, 0)
	if err = db.Where("student_id = ?", studentId).Preload("Task").Preload("Task.Subject").Preload("Task.Subject.Teacher").Find(&studentTasks).Error; err != nil {
		return
	}
	doneTaskIds := make([]uint, 0)
	for _, st := range studentTasks {
		doneTaskIds = append(doneTaskIds, st.TaskId)
	}

	tasks := make([]Task, 0)
	tx := db.Model(Task{})
	if len(doneTaskIds) > 0 {
		tx.Where("id NOT IN ?", doneTaskIds)
	}
	if err = tx.Preload("Subject").Preload("Subject.Teacher").Find(&tasks).Error; err != nil {
		return
	}
	res.Done = studentTasks
	for _, t := range tasks {
		if time.Time(t.Due).After(time.Now()) {
			res.Due = append(res.Due, t)
		} else {
			res.Late = append(res.Late, t)
		}
	}
	SortTasks(res)
	return
}

func SortTasks(tasks *Tasks) {
	sort.Slice(tasks.Done, func(i, j int) bool {
		return time.Time(tasks.Done[i].Task.Due).After(time.Time(tasks.Done[j].Task.Due))
	})
	sort.Slice(tasks.Late, func(i, j int) bool {
		return time.Time(tasks.Late[i].Due).After(time.Time(tasks.Late[j].Due))
	})
	sort.Slice(tasks.Due, func(i, j int) bool {
		return time.Time(tasks.Due[i].Due).After(time.Time(tasks.Due[j].Due))
	})
}
