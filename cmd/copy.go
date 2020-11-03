package cmd

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	// 等待组
	wgcp sync.WaitGroup

	// 上传本地文件到云端
	copyCmd = &cobra.Command{
		Use:     "cp object(s) ... bucket",
		Short:   "Copy local objects to a remote bucket",
		Long:    "Copy local objects to a remote bucket on MinIo/S3 Cloud Storage",
		Aliases: []string{"copy", "upload"},
		Example: "  Mydump2oss cp f1.sql f2.sql mysql_backup/",
		Run:     copyRun,
	}
)

func copyRun(cmd *cobra.Command, args []string) {
	length := len(args)
	if length < 2 {
		er(errors.New(cpUsage))
	}

	client := newClient()
	ctx := context.Background()
	bucket := trimSuffix(args[length-1], "/")
	checkBucket(client, ctx, bucket)

	for _, obj := range args[:length-1] {
		wgcp.Add(1)
		go cpObj(client, ctx, bucket, obj)
	}
	wgcp.Wait()
}

func cpObj(client *minio.Client, ctx context.Context, bucket, object string) {
	defer wgcp.Done()

	fileName := filepath.Base(object)

	putObjectOptions := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	_, err := client.FPutObject(ctx, bucket, fileName, object, putObjectOptions)
	if err != nil {
		er(err)
	} else {
		fmt.Println("Successfully uploaded object:", fileName)
	}
}
