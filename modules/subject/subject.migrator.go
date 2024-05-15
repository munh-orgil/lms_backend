package subject

import (
	"lms_backend/database"
	schedule_models "lms_backend/modules/schedule/models"
)

func RunMigrations() {
	db := database.DBconn

	// db.AutoMigrate(subject_models.Subject{})
	// db.AutoMigrate(subject_models.Task{})
	// db.AutoMigrate(subject_models.StudentSubject{})
	// db.AutoMigrate(subject_models.StudentTask{})
	// db.AutoMigrate(subject_models.Lecture{})
	// db.AutoMigrate(subject_models.Exam{})
	db.AutoMigrate(schedule_models.Schedule{})
}
