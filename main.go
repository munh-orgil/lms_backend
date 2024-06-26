package main

import (
	"flag"
	"lms_backend/database"
	"lms_backend/fiber_server"
	"lms_backend/global"
	"lms_backend/session"
)

var (
	migrate *bool
)

func main() {

	if err := global.LoadConfig("."); err != nil {
		panic(err.Error())
	}
	database.InitPostgres()

	migrate = flag.Bool("migrate", false, "a bool")
	flag.Parse()

	if *migrate {
		RunMigrations()
	}

	session.InitGjwt(global.Conf.JwtSecretPrvKeyPath, global.Conf.JwtSecretPubKeyPath)

	fiber_server.InitFiber()
}
