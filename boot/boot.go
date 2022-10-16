package boot

func Initialize() {
	initViper()
	initDB()
	initSession()
}
