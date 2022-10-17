package boot

func Initialize() {
	initViper()
	initLogger()
	initDB()
	initSession()
}
