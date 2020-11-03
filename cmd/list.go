package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

// 列出云端存储桶中文件
var listCmd = &cobra.Command{
	Use:     "ls bucket(s) ...",
	Short:   "List objects of remote bucket(s)",
	Long:    "List objects in remote bucket(s) on MinIo/S3 Cloud Storage",
	Aliases: []string{"list"},
	Example: "  Mydump2oss ls mysql_bucket/",
	Run:     listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		er(errors.New(lsUsage))
	}

	client := newClient()
	ctx := context.Background()
	listObjectsOptions := minio.ListObjectsOptions{Prefix: "", Recursive: true}
	for _, bucket := range args {
		bucket = trimSuffix(bucket, "/")
		checkBucket(ctx, client, bucket)

		objCh := client.ListObjects(ctx, bucket, listObjectsOptions)
		for obj := range objCh {
			if obj.Err != nil {
				er(obj.Err)
			}
			stringSize := size2string(float64(obj.Size))
			fmt.Println(obj.Key, stringSize)
		}
	}
}
