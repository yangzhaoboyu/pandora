package common

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"pandora/models"
	"path/filepath"
	"sync"
	"time"
)

//Upload files to Oss
func UploadOss(bucket *oss.Bucket, file string, wait *sync.WaitGroup) {
	wait.Add(1)
	_, fileName := filepath.Split(file)
	err := bucket.PutObjectFromFile(fileName, file)
	fmt.Printf("uploading %s ... \n ", file)
	if err != nil {
		fmt.Printf("%s upload failed:%s \n ", file, err)
	}
	record := &models.UploadRecord{}
	record.Source = file
	record.TargetDic = bucket.BucketName
	record.TargetSrc = fileName
	record.UpdateTimeStr = time.Now()
	_, err = models.AddRecord(record)
	if err != nil {
		fmt.Printf("%s upload failed:%s \n ", file, err)
	}
	wait.Done()
}
