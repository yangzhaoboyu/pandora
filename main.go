package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"pandora/common"
	"pandora/conf"
	"pandora/scanner"
	"path/filepath"
	"sync"
	"time"
)

var (
	config    conf.Config
	waitGroup sync.WaitGroup
)

func main() {

	confPath := flag.String("conf", "app.toml", "config path")
	flag.Parse()

	if *confPath == "" {
		return
	}
	err := common.LoadConfig(*confPath, &config)
	if err != nil {
		return
	}

	//OSS initialization
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(config.Oss.Bucket)
	if err != nil {
		panic(err)
	}

	filesChan := make(chan string)
	go filepath.Walk(config.Backup.Path, scanner.Visit(filesChan))

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
