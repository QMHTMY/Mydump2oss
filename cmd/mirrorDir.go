package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	// 等待组
	wgmr sync.WaitGroup

	// 同步本地目录下所有文件到云端
	mirrorCmd = &cobra.Command{
		Use:     "mr localDir/ bucket/",
		Short:   "Synchronize local objects to a remote bucket",
		Long:    "Synchronize local objects to a remote bucket on MinIo/S3 Cloud Storage",
		Aliases: []string{"mirror"},
		Example: "  Mydump2oss mr sql_bkp/ sql_bucket/",
		Run:     mirrorRun,
	}
)

func mirrorRun(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		er(errors.New(mrUsage))
	}

	// 获取localDir的reader
	separator := string(os.PathSeparator)
	localDir := trimSuffix(args[0], separator)
	reader, err := ioutil.ReadDir(localDir)
	if err != nil {
		er(err)
	}

	// 获取client并检查bucket
	client := newClient()
	ctx := context.Background()
	bucket := trimSuffix(args[1], separator)
	checkBucket(client, ctx, bucket)

	// 遍历并上传文件，若上传成功则删除本地文件
	fmt.Printf("Starting synchronize %s to bucket: %s\n", localDir, bucket)
	prePath := localDir + separator
	for _, file := range reader {
		if file.IsDir() {
			continue // 不递归上传子路径的内容
		}
		wgmr.Add(1)
		go putObject(client, ctx, prePath, file, bucket)
	}
	wgmr.Wait()
}

func putObject(client *minio.Client, ctx context.Context, pathStr string, file os.FileInfo, bucket string) {
	defer wgmr.Done()

	filePath := pathStr + file.Name()

	putObjOptions := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	_, err := client.FPutObject(ctx, bucket, file.Name(), filePath, putObjOptions)
	if err != nil {
		er(err)
	} else {
		fmt.Println("Successfully uploaded object:", filePath)
		if err := os.Remove(filePath); err != nil { // 上传后删除本地文件，节约空间
			er(err)
		} else {
			fmt.Println("Successfully deleted object:", filePath)
		}
	}
}
