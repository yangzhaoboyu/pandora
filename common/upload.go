package common

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path/filepath"
	"sync"
)

//Upload files to Oss
func UploadOss(bucket *oss.Bucket, file string, wait *sync.WaitGroup) {
	wait.Add(1)
	_, fileName := filepath.Split(file)
	err := bucket.PutObjectFromFile(fileName, file)
	fmt.Printf("uploading %s ...", file)
	if err != nil {
		fmt.Printf("%s upload failed:%s \n", file, err)
	}
	wait.Done()
}
