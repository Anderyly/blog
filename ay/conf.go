/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package ay

import (
	"blog/ay/lib"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	Yaml = getConfig()
}

func getConfig() *viper.Viper {
	dir := lib.NewDir()
	filePath, err := dir.LReadName("config.yaml", 3)
	if err != nil {
		log.Printf(err.Error())
	}
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to get the configuration.")
	}
	return viper.GetViper()
}

func WatchConf() {
	Yaml.WatchConfig()
	Yaml.OnConfigChange(func(event fsnotify.Event) {
		// 配置文件修改重新执行的方法
		Init()
		Logger.Info("Detect config change: " + event.String())
	})
}
