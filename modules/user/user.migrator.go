package user

import (
	"lms_backend/database"
	user_models "lms_backend/modules/user/models"
)

func RunMigrations() {
	db := database.DBconn

	db.AutoMigrate(user_models.User{})
}
