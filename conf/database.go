package conf

import "go-tagle/pkg/config"

var DBConf = struct {
	DBName string
}{}

func initDBConf() {
	DBConf.DBName = config.LoadString("database.dbname", "tagle.db")
}
