package boot

import "go-tagle/pkg/logger"

func initLogger() {
	logger.InitLogger("storage/logs/logs.log", 64, 7, 30, false, "single", "debug")
}
