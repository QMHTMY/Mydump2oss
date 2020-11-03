package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

// 同步本地目录下所有文件到云端
var mirrorCmd = &cobra.Command{
	Use:     "mr localDir/ bucket/",
	Short:   "Synchronize local objects to a remote bucket",
	Long:    "Synchronize local objects to a remote bucket on MinIo/S3 Cloud Storage",
	Aliases: []string{"mirror"},
	Example: "  Mydump2oss mr sql_bkp/ sql_bucket/",
	Run:     mirrorRun,
}

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

	// 获取client
	bucket := trimSuffix(args[1], separator)
	client := newClient()
	ctx := context.Background()
	checkBucket(ctx, client, bucket)

	// 遍历并上传文件，若上传成功则删除本地文件
	fmt.Printf("Starting synchronize %s to bucket:%s\n", localDir, bucket)
	putObjOptions := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	for _, file := range reader {
		if file.IsDir() {
			// 不递归上传子路径的内容
			continue
		}

		filePath := localDir + separator + file.Name()
		_, err := client.FPutObject(ctx, bucket, file.Name(), filePath, putObjOptions)
		if err != nil {
			er(err)
		} else {
			fmt.Println("Successfully uploaded object:", filePath)

			// 上传后删除本地文件，节约空间
			if err := os.Remove(filePath); err != nil {
				er(err)
			} else {
				fmt.Println("Successfully deleted object:", filePath)
			}
		}
	}
}
