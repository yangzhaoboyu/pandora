package conf

import "github.com/koding/multiconfig"

var AppConfig = &Config{}

//Configuration
type Config struct {
	Database DataBaseConfig
	Backup   BackupConfig
	Oss      OssConfig
}

type DataBaseConfig struct {
	Host     string
	Port     int
	DataBase string
	User     string
	PassWord string
}

type BackupConfig struct {
	Path string
}

type OssConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

//Load configuration file
func init() {
	loader := multiconfig.NewWithPath("app.toml")
	err := loader.Load(&AppConfig)
	if err != nil {
		panic(err)
	}
}
