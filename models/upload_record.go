package models

import "time"

//文件上传记录
type UploadRecord struct {
	Id            uint      `gorm:"primary_key;column:id"`
	Source        string    `gorm:"column:source"`
	TargetDic     string    `gorm:"column:target_dic"`
	TargetSrc     string    `gorm:"column:target_src"`
	UpdateTimeStr time.Time `gorm:"column:upload_time"`
	ThumbnailIUrl string    `gorm:"column:thumbnail_url"`
}

//自定义表名
func (UploadRecord) TableName() string {
	return "upload_record"
}

//添加上传记录
func AddRecord(record *UploadRecord) (*UploadRecord, error) {
	err := DB.Create(record).Error
	return record, err
}
