package conf

import "go-tagle/pkg/config"

var LogConf = struct {
	Filename   string
	MaxSize    int
	MaxBackup  int
	MaxAge     int
	IsCompress bool
	LogType    string
	Level      string
}{}

func initLogConf() {
	LogConf.Filename = config.LoadString("log.filename", "storage/logs/logs.log")
	LogConf.MaxSize = config.LoadInt("log.maxSize", 64)
	LogConf.MaxBackup = config.LoadInt("log.maxBackup", 7)
	LogConf.MaxAge = config.LoadInt("log.maxAge", 30)
	LogConf.IsCompress = config.LoadBool("log.IsCompress", false)
	LogConf.LogType = config.LoadString("log.logType", "daily")
	LogConf.Level = config.LoadString("log.level", "debug")
}
