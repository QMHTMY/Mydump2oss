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

// 同步本地目录下所有文件到minio/S3 Cloud Storage
// 使用方式： Mydump2oss mr localPath bucket
var mirrorCmd = &cobra.Command{
	Use:     "mr localdir bucket",
	Short:   "Synchronize local objects to a remote bucket",
	Long:    "Synchronize local objects to a remote bucket on MinIo/S3 Cloud Storage",
	Aliases: []string{"mirror"},
	Example: "  Mydump2oss mr sql_bkp sql_bucket",
	Run:     mirrorRun,
}

func mirrorRun(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		err := errors.New("Usage: Mydump2oss mr localdir bucket")
		er(err)
	}

	// 获取localDir的reader
	localDir := args[0]
	reader, err := ioutil.ReadDir(localDir)
	if err != nil {
		er(err)
	}

	bucket := args[1]
	client := newClient()
	ctx := context.Background()
	checkBucket(ctx, client, bucket)

	putObjOptions := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	for _, file := range reader {
		if file.IsDir() {
			continue // 不递归上传子路径的内容
		}

		filePath := localDir + string(os.PathSeparator) + file.Name()
		_, err := client.FPutObject(ctx, bucket, file.Name(), filePath, putObjOptions)
		if err != nil {
			er(err)
		} else {
			fmt.Println("Successfully uploaded object:", file.Name())
		}
	}
}
