package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

// 创建bucket 使用方式 Mydump2oss mb bucket1 bucket2 ...
var makeBucketCmd = &cobra.Command{
	Use:     "mb bucket(s) ...",
	Short:   "Make remote bucket(s)",
	Long:    "Make remote bucket(s) on MinIo/S3 Cloud Storage",
	Aliases: []string{"mkb", "mkbucket", "makebucket"},
	Example: "  Mydump2oss mb bucket1 bucket2 bucket3",
	Run:     makeBucketRun,
}

func makeBucketRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		err := errors.New("Usage: Mydump2oss mb bucket(s)...")
		er(err)
	}

	client := newClient()
	ctx := context.Background()
	bucketOptions := minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true}
	for _, bucket := range args {
        bucket = trimSuffix(bucket, "/")
		if err := client.MakeBucket(ctx, bucket, bucketOptions); err != nil {
			er(err)
		} else {
			fmt.Println("Successfully created bucket:", bucket)
		}
	}
}
