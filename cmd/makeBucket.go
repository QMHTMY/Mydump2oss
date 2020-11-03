package cmd

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
)

var (
	// 等待组
	wgmb sync.WaitGroup

	// 在云端创建文件桶bucket
	makeBucketCmd = &cobra.Command{
		Use:     "mb bucket(s) ...",
		Short:   "Make remote bucket(s)",
		Long:    "Make remote bucket(s) on MinIo/S3 Cloud Storage",
		Aliases: []string{"mkb", "mkbucket", "makebucket"},
		Example: "  Mydump2oss mb bucket1/ bucket2/ ...",
		Run:     makeBucketRun,
	}
)

func makeBucketRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		er(errors.New(mbUsage))
	}

	client := newClient()
	ctx := context.Background()
	for _, bucket := range args {
		wgmb.Add(1)
		go makeBucket(client, ctx, bucket)
	}
	wgmb.Wait()
}

func makeBucket(client *minio.Client, ctx context.Context, bucket string) {
	defer wgmb.Done()

	bucket = trimSuffix(bucket, "/")

	bucketOptions := minio.MakeBucketOptions{Region: region, ObjectLocking: true}
	if err := client.MakeBucket(ctx, bucket, bucketOptions); err != nil {
		er(err)
	} else {
		fmt.Println("Successfully created bucket:", bucket)
	}
}
