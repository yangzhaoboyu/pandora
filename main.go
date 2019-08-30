package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"pandora/common"
	"pandora/conf"
	"pandora/scanner"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	//OSS initialization
	client, err := oss.New(conf.AppConfig.Oss.Endpoint, conf.AppConfig.Oss.AccessKeyId, conf.AppConfig.Oss.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(conf.AppConfig.Oss.Bucket)
	if err != nil {
		panic(err)
	}
	filesChan := make(chan string)
	go filepath.Walk(conf.AppConfig.Backup.Path, scanner.Visit(filesChan))
	var waitGroup sync.WaitGroup
	for {
		select {
		case path := <-filesChan:
			go common.UploadOss(bucket, path, &waitGroup)
		case <-time.After(3 * time.Second):
			fmt.Println("End of path reading")
			goto end
		}
	}

end:
	waitGroup.Wait()
	fmt.Println("End of upload")
}
