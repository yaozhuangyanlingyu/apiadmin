package setting

import (
	"encoding/json"

	"github.com/apiadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 初始化框架
func Init() (err error) {
	// 初始化日志
	err = InitLogger()
	if err != nil {
		return err
	}

	// 初始化model
	err = models.Init()
	if err != nil {
		return err
	}

	return nil
}

// 初始化日志
func InitLogger() error {
	// 配置信息
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("log.path")
	config["level"] = GetLogLevel(beego.AppConfig.String("log.level"))
	config["maxlines"] = 100
	config["separate"] = []string{"error"}
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// 记录日志
	logs.Async()
	logs.EnableFuncCallDepth(true)
	logs.SetLogger(logs.AdapterMultiFile, string(configStr))
	logs.Debug("Start apiadmin for beego.")
	return nil
}

// 日志级别
func GetLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
