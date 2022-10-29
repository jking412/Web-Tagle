package conf

import "go-tagle/pkg/config"

var ServerConf = struct {
	Port string
}{}

func initServerConf() {
	ServerConf.Port = config.LoadString("server.port", "8000")
}
