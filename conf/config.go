package conf

//Configuration
type Config struct {
	Database DataBaseConfig
	Backup   BackupConfig
	Oss      OssConfig
}

type DataBaseConfig struct {
	Path string
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
