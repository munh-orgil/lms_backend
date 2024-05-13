package main

import "lms_backend/modules/subject"

func RunMigrations() {
	// log.RunMigrations()
	// otp.RunMigrations()
	// user.RunMigrations()
	subject.RunMigrations()
	// attachment.RunMigrations()
}
