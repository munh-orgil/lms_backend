package main

import (
	"lms_backend/modules/log"
	"lms_backend/modules/otp"
	"lms_backend/modules/subject"
	"lms_backend/modules/user"
)

func RunMigrations() {
	log.RunMigrations()
	otp.RunMigrations()
	user.RunMigrations()
	subject.RunMigrations()
}
